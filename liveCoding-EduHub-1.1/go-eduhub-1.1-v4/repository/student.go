package repository

import (
	//"encoding/json"
	"net/http"

	"a21hc3NpZ25tZW50/model"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	Store(student *model.Student) error
	ResetStudentRepo()
	GetByID(id int) (*model.Student, error)
}

type studentRepository struct {
	students []model.Student
}

func NewStudentRepo() *studentRepository {
	return &studentRepository{}
}

func (s *studentRepository) FetchAll() ([]model.Student, error) {
	return s.students, nil
}

func (s *studentRepository) Store(student *model.Student) error {
	s.students = append(s.students, *student)
	return nil
}

func (s *studentRepository) ResetStudentRepo() {
	s.students = []model.Student{}
}

func (s *studentRepository) GetByID(id int) (*model.Student, error) {
	for _, student := range s.students {
		if student.ID == id {
			return &student, nil
		}
	}
	return nil, nil // Not found
}

func GetStudentByID(id int, studentRepo StudentRepository) (int, interface{}) {
	student, err := studentRepo.GetByID(id)
	if err != nil {
		return http.StatusInternalServerError, map[string]string{"error": err.Error()}
	}
	if student == nil {
		return http.StatusNotFound, map[string]string{"error": "Student not found"}
	}
	return http.StatusOK, student
}

func AddStudent(student *model.Student, studentRepo StudentRepository) (int, interface{}) {
	err := studentRepo.Store(student)
	if err != nil {
		return http.StatusInternalServerError, map[string]string{"error": err.Error()}
	}
	return http.StatusOK, map[string]string{"message": "add course success"}
}
