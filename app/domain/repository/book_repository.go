package repository

import (
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateAt(b entity.Book) (entity.Book, error)
	Update(b entity.Book) (entity.Book, error)
	Delete(b entity.Book) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(ID uint) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAt(b entity.Book) (entity.Book, error) {
	err := r.db.Debug().Create(&b).Error
	return b, err
}

func (r *repository) Update(b entity.Book) (entity.Book, error) {
	err := r.db.Debug().Save(&b).Error
	return b, err
}

func (r *repository) Delete(b entity.Book) (entity.Book, error) {
	err := r.db.Debug().Delete(&b).Error
	return b, err
}

func (r *repository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Debug().Find(&books).Error
	return books, err
}

func (r *repository) FindByID(ID uint) (entity.Book, error) {
	var book entity.Book
	err := r.db.Debug().First(&book, ID).Error
	return book, err
}
