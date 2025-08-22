package handlers

import (
	"net/http"
	"strconv"
	"time"

	"library-management-system/internal/models"
	"library-management-system/internal/repository"
	"library-management-system/internal/utils"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	loanRepo   *repository.LoanRepository
	bookRepo   *repository.BookRepository
	memberRepo *repository.MemberRepository
}

func NewLoanHandler() *LoanHandler {
	return &LoanHandler{
		loanRepo:   repository.NewLoanRepository(),
		bookRepo:   repository.NewBookRepository(),
		memberRepo: repository.NewMemberRepository(),
	}
}

type CreateLoanRequest struct {
	BookID   uint      `json:"book_id" binding:"required"`
	MemberID uint      `json:"member_id" binding:"required"`
	DueDate  time.Time `json:"due_date" binding:"required"`
	Notes    string    `json:"notes"`
}

func GetAllLoans(c *gin.Context) {
	handler := NewLoanHandler()
	
	loans, err := handler.loanRepo.GetAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch loans")
		return
	}

	utils.SuccessResponse(c, "Loans retrieved successfully", loans)
}

func GetLoanByID(c *gin.Context) {
	handler := NewLoanHandler()
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid loan ID")
		return
	}

	loan, err := handler.loanRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Loan not found")
		return
	}

	utils.SuccessResponse(c, "Loan retrieved successfully", loan)
}

func CreateLoan(c *gin.Context) {
	handler := NewLoanHandler()
	
	var req CreateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	book, err := handler.bookRepo.GetByID(req.BookID)
	if err != nil {
		utils.NotFoundResponse(c, "Book not found")
		return
	}

	if book.Available <= 0 {
		utils.ErrorResponse(c, http.StatusBadRequest, "Book is not available for loan")
		return
	}

	member, err := handler.memberRepo.GetByID(req.MemberID)
	if err != nil {
		utils.NotFoundResponse(c, "Member not found")
		return
	}

	if member.Status != "active" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Member is not active")
		return
	}

	if req.DueDate.Before(time.Now()) {
		utils.ValidationErrorResponse(c, "Due date must be in the future")
		return
	}

	loan := &models.Loan{
		BookID:   req.BookID,
		MemberID: req.MemberID,
		LoanDate: time.Now(),
		DueDate:  req.DueDate,
		Status:   "borrowed",
		Notes:    req.Notes,
	}

	if err := handler.loanRepo.Create(loan); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create loan")
		return
	}

	book.Available--
	if err := handler.bookRepo.Update(book); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book availability")
		return
	}

	utils.SuccessResponse(c, "Loan created successfully", loan)
}

func ReturnBook(c *gin.Context) {
	handler := NewLoanHandler()
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid loan ID")
		return
	}

	loan, err := handler.loanRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Loan not found")
		return
	}

	if loan.Status == "returned" {
		utils.ErrorResponse(c, http.StatusBadRequest, "Book is already returned")
		return
	}

	now := time.Now()
	var fine float64 = 0
	if now.After(loan.DueDate) {
		daysOverdue := int(now.Sub(loan.DueDate).Hours() / 24)
		fine = float64(daysOverdue) * 1000
	}

	loan.ReturnDate = &now
	loan.Status = "returned"
	loan.Fine = fine

	if err := handler.loanRepo.Update(loan); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update loan")
		return
	}

	book, err := handler.bookRepo.GetByID(loan.BookID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get book")
		return
	}

	book.Available++
	if err := handler.bookRepo.Update(book); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book availability")
		return
	}

	utils.SuccessResponse(c, "Book returned successfully", loan)
}
