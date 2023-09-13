package helper

import (
	"github.com/raafly/api-classmate/entity"
	"github.com/raafly/api-classmate/model"
)

func ToStudentResponse(student entity.Student) model.StudentCreateResponse {
	return model.StudentCreateResponse{
		Nama: student.Nama,
		Absen: student.Absen,
		Gender: student.Gender,
		Nis: student.Nis,
	}
}

func ToStudentResponses(students []entity.Student) []model.StudentCreateResponse {
	var studentResponses []model.StudentCreateResponse 
	for _, student := range students {
		studentResponses = append(studentResponses, ToStudentResponse(student))
	}
	return studentResponses
}