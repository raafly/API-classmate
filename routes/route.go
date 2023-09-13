package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/api-classmate/controller"
)

func NewRouter(studentController controller.StudentController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/students", studentController.Insert)
	router.GET("/api/students", studentController.FindAll)
	router.GET("/api/students/:studentsId", studentController.FindById)
	router.POST("/api/students/:studentsId", studentController.Delete)

	return router
}