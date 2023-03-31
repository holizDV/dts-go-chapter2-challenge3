package main

import (
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/repository"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/domain/service"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/infrastructure/rest"
	"github.com/holizDV/dts-go-chapter2-challenge3/app/interfaces"
	"github.com/holizDV/dts-go-chapter2-challenge3/pkg/config"
)

func main() {

	dbConfig := config.PostgreConfig()
	newDb := rest.InitDatabase(dbConfig)

	bookRepository := repository.NewBookRepository(newDb)
	bookService := service.NewBookService(bookRepository)
	bookController := interfaces.NewBookController(bookService)

	router := config.NewRouter(bookController)

	router.Run(":8080")
}
