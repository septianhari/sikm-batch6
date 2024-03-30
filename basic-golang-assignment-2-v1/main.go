package main

import (
	"fmt"
	"strings"
)

// Define slice Students
var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

// Define map StudentStudyPrograms
var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"MI": "Manajemen Informasi",
	"GA": "Grafika Animasi",
	"SI": "Sistem Informasi",
}

// Define type for studentModifier function
type studentModifier func(string, *string)

// Function to update study program
func UpdateStudyProgram(programStudi string, student *string) {
	*student = (*student)[:len(*student)-2] + programStudi
}

// Function to perform student login
func Login(id, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}
	for _, student := range Students {
		data := splitStudentData(student)
		if data[0] == id && data[1] == name {
			return fmt.Sprintf("Login berhasil: %s", name)
		}
	}
	return "Login gagal: data mahasiswa tidak ditemukan"
}

// Function to register a new student
func Register(id, name, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}
	for _, student := range Students {
		data := splitStudentData(student)
		if data[0] == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}
	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

// Function to get study program by code
func GetStudyProgram(code string) string {
	program, found := StudentStudyPrograms[code]
	if !found {
		return "Kode program studi tidak ditemukan"
	}
	return program
}

// Function to modify student data
func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	for i, student := range Students {
		data := splitStudentData(student)
		if data[1] == nama {
			fn(programStudi, &Students[i])
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan."
}

// Helper function to split student data
func splitStudentData(student string) []string {
	return strings.Split(student, "_")
}

func main() {
	// Test cases
	fmt.Println(Login("", ""))             // Test Case 1
	fmt.Println(Login("A1234", "Aditira")) // Test Case 2

	fmt.Println(Register("A1234", "Aditira", "TI")) // Test Case 3
	fmt.Println(Login("A1234", "Aditira"))          // Test Case 3

	fmt.Println(GetStudyProgram("TI")) // Test Case 4

	fmt.Println(ModifyStudent("SI", "Afis", UpdateStudyProgram)) // Test Case 5
	fmt.Println(GetStudyProgram("MI"))                           // Test Case 5
}
