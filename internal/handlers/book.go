package handlers

import (
	"net/http"
	"strconv"

	"library-management-system/internal/models"
	"library-management-system/internal/repository"
	"library-management-system/internal/utils"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookRepo *repository.BookRepository
}

func NewBookHandler() *BookHandler {
	return &BookHandler{
		bookRepo: repository.NewBookRepository(),
	}
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Stock       int    `json:"stock" binding:"required,min=0"`
}

type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Stock       int    `json:"stock" binding:"min=0"`
}

func GetAllBooks(c *gin.Context) {
	handler := NewBookHandler()
	
	books, err := handler.bookRepo.GetAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch books")
		return
	}

	utils.SuccessResponse(c, "Books retrieved successfully", books)
}

func GetBookByID(c *gin.Context) {
	handler := NewBookHandler()
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid book ID")
		return
	}

	book, err := handler.bookRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Book not found")
		return
	}

	utils.SuccessResponse(c, "Book retrieved successfully", book)
}

func CreateBook(c *gin.Context) {
	handler := NewBookHandler()
	
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	existingBook, _ := handler.bookRepo.GetByISBN(req.ISBN)
	if existingBook != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Book with this ISBN already exists")
		return
	}

	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		Publisher:   req.Publisher,
		Year:        req.Year,
		Category:    req.Category,
		Description: req.Description,
		Stock:       req.Stock,
		Available:   req.Stock,
	}

	if err := handler.bookRepo.Create(book); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.SuccessResponse(c, "Book created successfully", book)
}

func UpdateBook(c *gin.Context) {
	handler := NewBookHandler()
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid book ID")
		return
	}

	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	book, err := handler.bookRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Book not found")
		return
	}

	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.ISBN != "" {
		if req.ISBN != book.ISBN {
			existingBook, _ := handler.bookRepo.GetByISBN(req.ISBN)
			if existingBook != nil {
				utils.ErrorResponse(c, http.StatusConflict, "Book with this ISBN already exists")
				return
			}
		}
		book.ISBN = req.ISBN
	}
	if req.Publisher != "" {
		book.Publisher = req.Publisher
	}
	if req.Year > 0 {
		book.Year = req.Year
	}
	if req.Category != "" {
		book.Category = req.Category
	}
	if req.Description != "" {
		book.Description = req.Description
	}
	if req.Stock >= 0 {
		stockDiff := req.Stock - book.Stock
		book.Stock = req.Stock
		book.Available += stockDiff
		if book.Available < 0 {
			book.Available = 0
		}
	}

	if err := handler.bookRepo.Update(book); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update book")
		return
	}

	utils.SuccessResponse(c, "Book updated successfully", book)
}

func DeleteBook(c *gin.Context) {
	handler := NewBookHandler()
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid book ID")
		return
	}

	_, err = handler.bookRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Book not found")
		return
	}

	if err := handler.bookRepo.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	utils.SuccessResponse(c, "Book deleted successfully", nil)
}
