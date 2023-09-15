package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/raafly/api-classmate/helper"
	"github.com/raafly/api-classmate/model"
	"github.com/raafly/api-classmate/service"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentControllerImpl(studentService service.StudentService) StudentControllerImpl {
	return StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) Insert(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentCreateRequest := model.StudentCreateRequest{}
	helper.ReadFromRequestBody(r, &studentCreateRequest)

	student := controller.StudentService.Insert(r.Context(), studentCreateRequest)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "OK",
		Data: student,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentResponse := controller.StudentService.FindAll(r.Context())
	webResponse := model.WebResponse {
		Code: 201,
		Status: "OK",
		Data: studentResponse,
	}

	 helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentsId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	studentResponse := controller.StudentService.FindById(r.Context(), id)
	webResponse := model.WebResponse {
		Code: 201,
		Status: "OK",
		Data: studentResponse,
	}

	 helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentId := params.ByName("studentsId")
	id, err := strconv.Atoi(studentId)
	helper.PanicIfError(err)

	controller.StudentService.Delete(r.Context(), id)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}
