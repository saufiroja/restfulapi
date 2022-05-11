package main

import (
	"net/http"
	"restapi-golang/app"
	"restapi-golang/controllers"
	"restapi-golang/exception"
	"restapi-golang/helper"
	"restapi-golang/middlewares"
	"restapi-golang/repository"
	"restapi-golang/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryControllers := controllers.NewCategoryControllers(categoryService)

	router := httprouter.New()

	router.GET("/api/category", categoryControllers.FindAll)
	router.GET("/api/category/:categoryId", categoryControllers.FindById)
	router.POST("/api/category", categoryControllers.Create)
	router.PUT("/api/category/:categoryId", categoryControllers.Update)
	router.DELETE("/api/category/:categoryId", categoryControllers.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middlewares.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
