package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/api-classmate/entity"
	"github.com/raafly/api-classmate/helper"
)

type studentRepositoryImpl struct {
}

func NewStudentRepositoryImpl() StudentRepository {
	return &studentRepositoryImpl{}
}

func (repository *studentRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, student entity.Student) entity.Student{
	script := "INSERT INTO students(nama, absen, gender, nis) VALUES(?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, script, student.Nama, student.Absen, student.Gender, student.Nis)
	helper.PanicIfError(err)

	return student
}

func (repository *studentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Student {
	script := "SELECT nama, absen, gender, nis FROM students"
	rows, err := tx.QueryContext(ctx, script)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []entity.Student
	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Nama, &student.Absen, &student.Gender, &student.Nis)
		helper.PanicIfError(err)
		students = append(students, student)

	}
	return students
} 


func (repository *studentRepositoryImpl) FindById(ctx context.Context , tx *sql.Tx, studentNis int) (entity.Student, error) {
	script := "SELECT nama, absen, gender, nis FROM students WHERE nis = ?"
	rows, err := tx.QueryContext(ctx, script, studentNis)
	helper.PanicIfError(err)
	defer rows.Close()

	student := entity.Student{}
	if rows.Next() {
		err := rows.Scan(&student.Nama, &student.Absen, &student.Gender, &student.Nis)
		helper.PanicIfError(err)
		return student, nil
	} else {
		return student, errors.New("student is not found")
	}
}

func (repository *studentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, studentNis int)  {
	script := "DELETE FROM students WHERE nis = ?"
	_, err := tx.ExecContext(ctx, script, studentNis)
	helper.PanicIfError(err)
}

