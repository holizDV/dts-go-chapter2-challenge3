package config

import (
	"github.com/gin-gonic/gin"

	"github.com/holizDV/dts-go-chapter2-challenge3/app/interfaces"
)

func NewRouter(c interfaces.BookController) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("product/v1")

	v1.POST("/books", c.CreateAt)
	v1.GET("/books", c.FindAll)
	v1.GET("/books/:bookID", c.FindByID)
	v1.DELETE("/books/:bookID", c.Delete)
	v1.PUT("/books/:bookID", c.Update)

	return router
}
