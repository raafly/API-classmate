package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/api-classmate/entity"
	"github.com/raafly/api-classmate/helper"
	"github.com/raafly/api-classmate/model"
	"github.com/raafly/api-classmate/repository"
	"github.com/raafly/api-classmate/exception"
)

type studentServiceImpl struct {
	StudentRepository repository.StudentRepository
	DB *sql.DB
	Validate validator.Validate
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, DB *sql.DB, validate *validator.Validate) *studentServiceImpl {
	return &studentServiceImpl{
		StudentRepository: studentRepository,
		DB: DB,
		Validate: *validator.New(),
	}
}

func (service *studentServiceImpl) Insert(ctx context.Context, request model.StudentCreateRequest) model.StudentCreateResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student := entity.Student {
		Nama: request.Nama,
		Absen: request.Absen,
		Gender: request.Gender,
		Nis: request.Nis,
	}

	student = service.StudentRepository.Insert(ctx, tx, student)
	return helper.ToStudentResponse(student)
}

func (service *studentServiceImpl) FindAll(ctx context.Context) []model.StudentCreateResponse  {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	students := service.StudentRepository.FindAll(ctx, tx)
	return helper.ToStudentResponses(students)
}

func (service *studentServiceImpl) FindById(ctx context.Context, studentNis int) model.StudentCreateResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FindById(ctx, tx, studentNis)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToStudentResponse(student)
}

func (service *studentServiceImpl) Delete(ctx context.Context, studentNis int)  {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FindById(ctx, tx, studentNis)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.StudentRepository.Delete(ctx, tx, student.Nis)
}	
