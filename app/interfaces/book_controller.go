package interfaces

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/entity"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/service"
	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/helper"
)

type BookController interface {
	CreateAt(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type controller struct {
	service service.BookService
}

func NewBookController(service service.BookService) *controller {
	return &controller{service}
}

func (c *controller) CreateAt(ctx *gin.Context) {
	var bookRequest entity.BookRequest
	var response entity.BaseHttpResponse

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorFieldResponse(err, http.StatusBadRequest))
		return
	}

	book, err := c.service.CreateAt(bookRequest)
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorMessageResponse(err, http.StatusBadRequest))
		return
	}

	helper.HttpRespoonse(ctx, response.SuccessResponse(convertToBookResponse(book)))
}

func (c *controller) Update(ctx *gin.Context) {
	var bookRequest entity.BookRequest
	var response entity.BaseHttpResponse

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorFieldResponse(err, http.StatusBadRequest))
		return
	}

	bookID := ctx.Params.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	helper.PanicIfError(err)

	book, err := c.service.Update(uint(id), bookRequest)
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorMessageResponse(err, http.StatusNotFound))
		return
	}

	helper.HttpRespoonse(ctx, response.SuccessResponse(convertToBookResponse(book)))
}

func (c *controller) Delete(ctx *gin.Context) {
	var response entity.BaseHttpResponse

	bookID := ctx.Params.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	helper.PanicIfError(err)

	book, err := c.service.Delete(uint(id))
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorMessageResponse(err, http.StatusNotFound))
		return
	}

	helper.HttpRespoonse(ctx, response.SuccessResponse(convertToBookDeleteResponse(book)))
}

func (c *controller) FindAll(ctx *gin.Context) {
	var response entity.BaseHttpResponse
	var bookResponse []entity.BookResponse

	books, err := c.service.FindAll()
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorMessageResponse(err, http.StatusBadRequest))
		return
	}

	for _, book := range books {
		response := convertToBookResponse(book)
		bookResponse = append(bookResponse, response)
	}

	helper.HttpRespoonse(ctx, response.SuccessResponse(bookResponse))
}

func (c *controller) FindByID(ctx *gin.Context) {
	var response entity.BaseHttpResponse

	bookID := ctx.Params.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	helper.PanicIfError(err)

	book, err := c.service.FindByID(uint(id))
	if err != nil {
		helper.HttpRespoonse(ctx, response.ErrorMessageResponse(err, http.StatusNotFound))
		return
	}
	helper.HttpRespoonse(ctx, response.SuccessResponse(convertToBookResponse(book)))
}

func convertToBookResponse(b entity.Book) entity.BookResponse {
	return entity.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}
}

func convertToBookDeleteResponse(b entity.Book) entity.BookDeleteResponse {
	return entity.BookDeleteResponse{
		ID: b.ID,
	}
}
