package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"library-management-system/internal/models"
	"library-management-system/internal/repository"
	"library-management-system/internal/utils"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	memberRepo *repository.MemberRepository
}

func NewMemberHandler() *MemberHandler {
	return &MemberHandler{
		memberRepo: repository.NewMemberRepository(),
	}
}

type CreateMemberRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateMemberRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

func generateMemberCode() string {
	return fmt.Sprintf("MEM%06d", time.Now().Unix()%1000000)
}

func GetAllMembers(c *gin.Context) {
	handler := NewMemberHandler()

	members, err := handler.memberRepo.GetAll()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch members")
		return
	}

	utils.SuccessResponse(c, "Members retrieved successfully", members)
}

func GetMemberByID(c *gin.Context) {
	handler := NewMemberHandler()

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid member ID")
		return
	}

	member, err := handler.memberRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Member not found")
		return
	}

	utils.SuccessResponse(c, "Member retrieved successfully", member)
}

func CreateMember(c *gin.Context) {
	handler := NewMemberHandler()

	var req CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	existingMember, _ := handler.memberRepo.GetByEmail(req.Email)
	if existingMember != nil {
		utils.ErrorResponse(c, http.StatusConflict, "Member with this email already exists")
		return
	}

	memberCode := generateMemberCode()
	for {
		existingCode, _ := handler.memberRepo.GetByMemberCode(memberCode)
		if existingCode == nil {
			break
		}
		memberCode = generateMemberCode()
	}

	member := &models.Member{
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		MemberCode: memberCode,
		Status:     "active",
	}

	if err := handler.memberRepo.Create(member); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create member")
		return
	}

	utils.SuccessResponse(c, "Member created successfully", member)
}

func UpdateMember(c *gin.Context) {
	handler := NewMemberHandler()

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid member ID")
		return
	}

	var req UpdateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, err.Error())
		return
	}

	member, err := handler.memberRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Member not found")
		return
	}

	if req.Name != "" {
		member.Name = req.Name
	}
	if req.Email != "" {
		if req.Email != member.Email {
			existingMember, _ := handler.memberRepo.GetByEmail(req.Email)
			if existingMember != nil {
				utils.ErrorResponse(c, http.StatusConflict, "Member with this email already exists")
				return
			}
		}
		member.Email = req.Email
	}
	if req.Phone != "" {
		member.Phone = req.Phone
	}
	if req.Address != "" {
		member.Address = req.Address
	}
	if req.Status != "" {
		member.Status = req.Status
	}

	if err := handler.memberRepo.Update(member); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update member")
		return
	}

	utils.SuccessResponse(c, "Member updated successfully", member)
}

func DeleteMember(c *gin.Context) {
	handler := NewMemberHandler()

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ValidationErrorResponse(c, "Invalid member ID")
		return
	}

	_, err = handler.memberRepo.GetByID(uint(id))
	if err != nil {
		utils.NotFoundResponse(c, "Member not found")
		return
	}

	if err := handler.memberRepo.Delete(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete member")
		return
	}

	utils.SuccessResponse(c, "Member deleted successfully", nil)
}
