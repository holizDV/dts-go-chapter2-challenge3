package service

import (
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/entity"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/repository"
	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/helper"
)

type BookService interface {
	CreateAt(b entity.BookRequest) (entity.Book, error)
	Update(ID uint, b entity.BookRequest) (entity.Book, error)
	Delete(ID uint) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(ID uint) (entity.Book, error)
}

type service struct {
	repository repository.BookRepository
}

func NewBookService(r repository.BookRepository) *service {
	return &service{repository: r}
}

func (s *service) CreateAt(b entity.BookRequest) (entity.Book, error) {
	book := entity.Book{
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}
	return s.repository.CreateAt(book)
}

func (s *service) Update(ID uint, b entity.BookRequest) (entity.Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		helper.PanicIfError(err)
	}

	book.Title = b.Title
	book.Author = b.Author
	book.Description = b.Description

	return s.repository.Update(book)
}

func (s *service) Delete(ID uint) (entity.Book, error) {
	book, _ := s.repository.FindByID(ID)
	return s.repository.Delete(book)
}

func (s *service) FindAll() ([]entity.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID uint) (entity.Book, error) {
	return s.repository.FindByID(ID)
}
