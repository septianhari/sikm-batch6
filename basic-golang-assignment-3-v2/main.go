package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID           string
	Name         string
	StudyProgram string
}

type StudentManager interface {
	GetStudents() []Student
	Login(id, name string) (string, error)
	Register(id, name, studyProgram string) (string, error)
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn StudentModifier) (string, error)
}

type InMemoryStudentManager struct {
	students             []Student
	studentStudyPrograms map[string]string
}

type StudentModifier func(*Student) error

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []Student{
			{ID: "A12345", Name: "Aditira", StudyProgram: "TI"},
			{ID: "B21313", Name: "Dito", StudyProgram: "TK"},
			{ID: "A34555", Name: "Afis", StudyProgram: "MI"},
		},
		studentStudyPrograms: map[string]string{
			"SI": "Sistem Informasi",
			"TK": "Teknik Komputer",
			"TI": "Teknik Informatika",
			"MI": "Manajemen Informatika",
		},
	}
}

func (ism *InMemoryStudentManager) GetStudents() []Student {
	return ism.students
}

func (ism *InMemoryStudentManager) Login(id, name string) (string, error) {
	if id == "" || name == "" {
		return "", errors.New("ID or Name is undefined!")
	}

	for _, student := range ism.students {
		if student.ID == id && student.Name == name {
			return fmt.Sprintf("Login berhasil: %s", name), nil
		}
	}
	return "", errors.New("Login gagal: data mahasiswa tidak ditemukan")
}

func (ism *InMemoryStudentManager) Register(id, name, studyProgram string) (string, error) {
	if id == "" || name == "" || studyProgram == "" {
		return "", errors.New("ID, Name or StudyProgram is undefined!")
	}

	if _, exists := ism.studentStudyPrograms[studyProgram]; !exists {
		return "", errors.New(fmt.Sprintf("Study program %s is not found", studyProgram))
	}

	for _, student := range ism.students {
		if student.ID == id {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		}
	}

	ism.students = append(ism.students, Student{ID: id, Name: name, StudyProgram: studyProgram})
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil
}

func (ism *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {
		return "", errors.New("Code is undefined!")
	}

	program, exists := ism.studentStudyPrograms[code]
	if !exists {
		return "", errors.New("Kode program studi tidak ditemukan")
	}

	return program, nil
}

func (ism *InMemoryStudentManager) ModifyStudent(name string, fn StudentModifier) (string, error) {
	if name == "" {
		return "", errors.New("Mahasiswa tidak ditemukan.")
	}

	for i, student := range ism.students {
		if student.Name == name {
			if err := fn(&ism.students[i]); err != nil {
				return "", err
			}
			return "Program studi mahasiswa berhasil diubah.", nil
		}
	}

	return "", errors.New("Mahasiswa tidak ditemukan.")
}

func (ism *InMemoryStudentManager) ChangeStudyProgram(programStudi string) StudentModifier {
	return func(student *Student) error {
		if _, exists := ism.studentStudyPrograms[programStudi]; !exists {
			return errors.New("Kode program studi tidak ditemukan")
		}
		student.StudyProgram = programStudi
		return nil
	}
}
