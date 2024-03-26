package main

import "fmt"

// Students berisi data mahasiswa dengan format "id_nama_jurusan"
var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

// StudentStudyPrograms memetakan kode program studi ke nama program studi
var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"MI": "Manajemen Informasi",
}

// Login berfungsi untuk melakukan login dengan memeriksa keberadaan data mahasiswa
func Login(id, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	for _, student := range Students {
		data := parseStudentData(student)
		if data.id == id && data.name == name {
			return fmt.Sprintf("Login berhasil: %s", name)
		}
	}
	return "Login gagal: data mahasiswa tidak ditemukan"
}

// Register berfungsi untuk mendaftarkan mahasiswa baru
func Register(id, name, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}

	for _, student := range Students {
		data := parseStudentData(student)
		if data.id == id {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

// GetStudyProgram berfungsi untuk mendapatkan nama program studi berdasarkan kode program studi
func GetStudyProgram(code string) string {
	if program, ok := StudentStudyPrograms[code]; ok {
		return program
	}
	return "Kode program studi tidak ditemukan"
}

// ModifyStudent berfungsi untuk mengubah program studi mahasiswa
func ModifyStudent(programStudi, nama string, fn func(string, *string)) string {
	for i, student := range Students {
		data := parseStudentData(student)
		if data.name == nama {
			fn(programStudi, &Students[i])
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan."
}

// UpdateStudyProgram adalah fungsi yang akan dijalankan pada fungsi ModifyStudent
func UpdateStudyProgram(programStudi string, student *string) {
	*student = fmt.Sprintf("%s_%s", (*student)[:len(*student)-2], programStudi)
}

// studentData adalah struct untuk menyimpan data mahasiswa
type studentData struct {
	id    string
	name  string
	major string
}

// parseStudentData berfungsi untuk memecah data mahasiswa menjadi id, name, dan major
func parseStudentData(student string) studentData {
	var id, name, major string
	fmt.Sscanf(student, "%s_%s_%s", &id, &name, &major)
	return studentData{id, name, major}
}

func main() {
	fmt.Println(Login("", ""))                                   // Test Case 1
	fmt.Println(Login("A1234", "Aditira"))                       // Test Case 2
	fmt.Println(Register("A1234", "Aditira", "TI"))              // Test Case 3
	fmt.Println(Login("A1234", "Aditira"))                       // Test Case 3
	fmt.Println(GetStudyProgram("TI"))                           // Test Case 4
	fmt.Println(ModifyStudent("SI", "Afis", UpdateStudyProgram)) // Test Case 5
	fmt.Println(GetStudyProgram("MI"))                           // Test Case 5
}
