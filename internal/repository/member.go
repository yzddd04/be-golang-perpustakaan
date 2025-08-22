package repository

import (
	"library-management-system/internal/config"
	"library-management-system/internal/models"
)

type MemberRepository struct{}

func NewMemberRepository() *MemberRepository {
	return &MemberRepository{}
}

func (r *MemberRepository) Create(member *models.Member) error {
	return config.GetDB().Create(member).Error
}

func (r *MemberRepository) GetAll() ([]models.Member, error) {
	var members []models.Member
	err := config.GetDB().Find(&members).Error
	return members, err
}

func (r *MemberRepository) GetByID(id uint) (*models.Member, error) {
	var member models.Member
	err := config.GetDB().First(&member, id).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepository) Update(member *models.Member) error {
	return config.GetDB().Save(member).Error
}

func (r *MemberRepository) Delete(id uint) error {
	return config.GetDB().Delete(&models.Member{}, id).Error
}

func (r *MemberRepository) GetByEmail(email string) (*models.Member, error) {
	var member models.Member
	err := config.GetDB().Where("email = ?", email).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepository) GetByMemberCode(memberCode string) (*models.Member, error) {
	var member models.Member
	err := config.GetDB().Where("member_code = ?", memberCode).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}
