package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-playground/validator/v10"
	"github.com/raafly/api-classmate/config"
	"github.com/raafly/api-classmate/controller"
	"github.com/raafly/api-classmate/helper"
	"github.com/raafly/api-classmate/repository"
	"github.com/raafly/api-classmate/routes"
	"github.com/raafly/api-classmate/service"
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
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)


}