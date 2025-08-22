package repository

import (
	"library-management-system/internal/config"
	"library-management-system/internal/models"
)

type LoanRepository struct{}

func NewLoanRepository() *LoanRepository {
	return &LoanRepository{}
}

func (r *LoanRepository) Create(loan *models.Loan) error {
	return config.GetDB().Create(loan).Error
}

func (r *LoanRepository) GetAll() ([]models.Loan, error) {
	var loans []models.Loan
	err := config.GetDB().Preload("Book").Preload("Member").Find(&loans).Error
	return loans, err
}

func (r *LoanRepository) GetByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	err := config.GetDB().Preload("Book").Preload("Member").First(&loan, id).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *LoanRepository) Update(loan *models.Loan) error {
	return config.GetDB().Save(loan).Error
}

func (r *LoanRepository) GetByMemberID(memberID uint) ([]models.Loan, error) {
	var loans []models.Loan
	err := config.GetDB().Preload("Book").Where("member_id = ?", memberID).Find(&loans).Error
	return loans, err
}

func (r *LoanRepository) GetByBookID(bookID uint) ([]models.Loan, error) {
	var loans []models.Loan
	err := config.GetDB().Preload("Member").Where("book_id = ?", bookID).Find(&loans).Error
	return loans, err
}

func (r *LoanRepository) GetActiveLoans() ([]models.Loan, error) {
	var loans []models.Loan
	err := config.GetDB().Preload("Book").Preload("Member").Where("status = ?", "borrowed").Find(&loans).Error
	return loans, err
}

func (r *LoanRepository) GetOverdueLoans() ([]models.Loan, error) {
	var loans []models.Loan
	err := config.GetDB().Preload("Book").Preload("Member").
		Where("status = ? AND due_date < NOW()", "borrowed").
		Find(&loans).Error
	return loans, err
}
