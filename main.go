package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/raafly/api-classmate/config"
	"github.com/raafly/api-classmate/controller"
	"github.com/raafly/api-classmate/helper"
	"github.com/raafly/api-classmate/repository"
	"github.com/raafly/api-classmate/routes"
	"github.com/raafly/api-classmate/service"
	"github.com/raafly/api-classmate/middleware"
)

func main() {

	DB := config.NewDB()
	validate := validator.New()
	studentRepository := repository.NewStudentRepositoryImpl()
	studentService := service.NewStudentServiceImpl(studentRepository, DB, validate)
	studentController := controller.NewStudentControllerImpl(studentService)
	router := routes.NewRouter(&studentController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)


}