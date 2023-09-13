package repository

import (
	"context"
	"database/sql"

	"github.com/raafly/api-classmate/entity"
)

type StudentRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, student entity.Student) entity.Student
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Student
	FindById(ctx context.Context , tx *sql.Tx, studentNis int) (entity.Student, error)
	Delete(ctx context.Context, tx *sql.Tx, studentNis int) 
}