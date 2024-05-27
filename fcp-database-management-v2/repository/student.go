package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	result := s.db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}
	return students, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	result := s.db.Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	result := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	result := s.db.Delete(&model.Student{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	result := s.db.First(&student, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var studentClasses []model.StudentClass
	result := s.db.Table("students").
		Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
		Joins("LEFT JOIN classes ON students.class_id = classes.id").
		Find(&studentClasses)

	if result.Error != nil {
		return nil, result.Error
	}

	return &studentClasses, nil // TODO: replace this
}
