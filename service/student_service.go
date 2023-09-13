package service

import (
	"context"
	"github.com/raafly/api-classmate/model"
)

type StudentService interface {
	Insert(ctx context.Context, request model.StudentCreateRequest) model.StudentCreateResponse
	FindAll(ctx context.Context) []model.StudentCreateResponse 
	FindById(ctx context.Context, studentNis int) model.StudentCreateResponse
	Delete(ctx context.Context, studentNis int) 
}