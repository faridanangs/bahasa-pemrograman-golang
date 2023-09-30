package main

import (
	"golangapi/app"
	"golangapi/controller"
	"golangapi/helper"
	"golangapi/middleware"
	"golangapi/repository"
	"golangapi/service"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func TestMain(t *testing.T) {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryConntroller := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryConntroller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
