package app

import (
	"yudistirachma/go-restful-api/controller"
	"yudistirachma/go-restful-api/exception"
	"yudistirachma/go-restful-api/repository"
	"yudistirachma/go-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	db := NewDB()
	validate := validator.New()

	// category
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)
	categoryController := controller.NewCategoryController(service.NewCategoryService(repository.NewCategoryRepository(), db, validate))

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
