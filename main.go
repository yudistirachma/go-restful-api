package main

import (
	"net/http"
	"yudistirachma/go-restful-api/app"
	"yudistirachma/go-restful-api/controller"
	"yudistirachma/go-restful-api/helper"
	"yudistirachma/go-restful-api/middleware"
	"yudistirachma/go-restful-api/repository"
	"yudistirachma/go-restful-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
