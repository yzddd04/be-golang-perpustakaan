package repository

import (
	"library-management-system/internal/config"
	"library-management-system/internal/models"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(book *models.Book) error {
	return config.GetDB().Create(book).Error
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := config.GetDB().Find(&books).Error
	return books, err
}

func (r *BookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := config.GetDB().First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	return config.GetDB().Save(book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return config.GetDB().Delete(&models.Book{}, id).Error
}

func (r *BookRepository) GetByISBN(isbn string) (*models.Book, error) {
	var book models.Book
	err := config.GetDB().Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) UpdateStock(id uint, available int) error {
	return config.GetDB().Model(&models.Book{}).Where("id = ?", id).Update("available", available).Error
}
