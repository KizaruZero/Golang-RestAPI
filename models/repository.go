package models

import "gorm.io/gorm"

type Repository interface {
	GetBooks() ([]Book, error)
	GetBookByID(id int) (Book, error)
	CreateBook(book Book) (Book, error)
	UpdateBook(id int, book Book) (Book, error)
	DeleteBook(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetBooks() ([]Book, error) {
	var books []Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *repository) GetBookByID(id int) (Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) CreateBook(book Book) (Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) UpdateBook(id int, book Book) (Book, error) {
	if err := r.db.Model(&Book{}).Where("id = ?", id).Updates(book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) DeleteBook(id int) error {
	if err := r.db.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}