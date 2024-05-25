package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Query("SELECT * FROM students")
	if errors.Is(err, sql.ErrNoRows) {
		return []model.Student{}, errors.New("data tidak ditemukan")
	} else if err != nil {
		return []model.Student{}, err
	}
	defer rows.Close()

	students := []model.Student{}
	for rows.Next() {
		student := model.Student{}
		if err := rows.Scan(&student.ID, &student.Name, &student.Address, &student.Class); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	_, err := s.db.Exec("INSERT INTO students (name, address, class) VALUES ($1, $2, $3)", student.Name, student.Address, student.Class)
	if err != nil {
		return err
	}
	return nil
}
