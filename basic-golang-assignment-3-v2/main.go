package main

import (
	"errors"
	"fmt"
)

// StudentManager merupakan interface yang mendefinisikan fungsi-fungsi untuk student portal.
type StudentManager interface {
	GetStudents() []Student
	Login(id, name string) (string, error)
	Register(id, name, studyProgram string) (string, error)
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn StudentModifier) (string, error)
}

// StudentModifier merupakan tipe fungsi yang digunakan untuk memodifikasi data mahasiswa.
type StudentModifier func(student *Student) error

// Student adalah struktur data untuk merepresentasikan seorang mahasiswa.
type Student struct {
	ID           string
	Name         string
	StudyProgram string
}

// InMemoryStudentManager adalah implementasi dari StudentManager yang menggunakan penyimpanan dalam memori.
type InMemoryStudentManager struct {
	students             []Student
	studentStudyPrograms map[string]string
}

// NewInMemoryStudentManager membuat instans baru dari InMemoryStudentManager.
func NewInMemoryStudentManager() *InMemoryStudentManager {
	// Contoh inisialisasi map program studi mahasiswa
	studentStudyPrograms := map[string]string{
		"SI": "Sistem Informasi",
		"TK": "Teknik Komputer",
		"TI": "Teknik Informatika",
	}

	return &InMemoryStudentManager{
		students:             []Student{},
		studentStudyPrograms: studentStudyPrograms,
	}
}

// GetStudents mengembalikan seluruh data mahasiswa.
func (ism *InMemoryStudentManager) GetStudents() []Student {
	return ism.students
}

// Login memeriksa apakah data mahasiswa terdapat dalam slice students.
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

// Register menambahkan data mahasiswa baru ke dalam slice students.
func (ism *InMemoryStudentManager) Register(id, name, studyProgram string) (string, error) {
	if id == "" || name == "" || studyProgram == "" {
		return "", errors.New("ID, Name or StudyProgram is undefined!")
	}

	if _, ok := ism.studentStudyPrograms[studyProgram]; !ok {
		return "", fmt.Errorf("Study program %s is not found", studyProgram)
	}

	for _, student := range ism.students {
		if student.ID == id {
			return "", errors.New("Registrasi gagal: id sudah digunakan")
		}
	}

	ism.students = append(ism.students, Student{ID: id, Name: name, StudyProgram: studyProgram})
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, ism.studentStudyPrograms[studyProgram]), nil
}

// GetStudyProgram mencari nama program studi dari map studentStudyPrograms berdasarkan kode program studi yang dimasukkan.
func (ism *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {
		return "", errors.New("Code is undefined!")
	}

	if program, ok := ism.studentStudyPrograms[code]; ok {
		return program, nil
	}

	return "", errors.New("Kode program studi tidak ditemukan")
}

// ModifyStudent mencari data mahasiswa berdasarkan nama yang dimasukkan, dan akan memodifikasinya menggunakan fungsi fn.
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

// ChangeStudyProgram mengembalikan sebuah fungsi yang akan memodifikasi program studi dari mahasiswa yang diberikan.
func (ism *InMemoryStudentManager) ChangeStudyProgram(programStudi string) StudentModifier {
	return func(student *Student) error {
		if _, ok := ism.studentStudyPrograms[programStudi]; !ok {
			return errors.New("Kode program studi tidak ditemukan")
		}
		student.StudyProgram = programStudi
		return nil
	}
}
