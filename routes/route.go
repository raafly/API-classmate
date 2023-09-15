package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/api-classmate/controller"
	"github.com/raafly/api-classmate/exception"
)

func NewRouter(studentController controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/students", studentController.Insert)
	router.GET("/api/students", studentController.FindAll)
	router.GET("/api/students/:studentsId", studentController.FindById)
	router.DELETE("/api/students/:studentsId", studentController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}