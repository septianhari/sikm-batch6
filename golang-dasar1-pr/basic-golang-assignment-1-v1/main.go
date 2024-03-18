package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI, A6666_Hari_KOM"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi, KOM_Komedi"

func Login(id string, name string) string {
	var statusLogin string
	var students []string

	students = strings.Split(Students, ", ")

	if id == "" || name == "" {
		statusLogin = "ID or Name is undefined!"
	} else if len(id) != 5 {
		statusLogin = "ID must be 5 characters long!"
	} else {
		for i := 0; i < len(students); i++ {
			if strings.Split(students[i], "_")[0] == id && strings.Split(students[i], "_")[1] == name {
				statusLogin = "Login berhasil: " + name + " (" + strings.Split(students[i], "_")[2] + ")"
				break
			} else {
				statusLogin = "Login gagal: data mahasiswa tidak ditemukan"
			}
		}
	}
	return statusLogin // TODO: replace this
}

func Register(id string, name string, major string) string {
	var statusRegister string
	var students []string

	students = strings.Split(Students, ", ")

	if id == "" || name == "" || major == "" {
		statusRegister = "ID, Name or Major is undefined!"
	} else if len(id) != 5 {
		statusRegister = "ID must be 5 characters long!"
	} else {
		for i := 0; i < len(students); i++ {
			if strings.Split(students[i], "_")[0] == id {
				statusRegister = "Registrasi gagal: id sudah digunakan"
				break
			} else {
				statusRegister = "Registrasi berhasil: " + name + " (" + major + ")"
			}
		}
	}
	return statusRegister // TODO: SELESAI REGISTER
}

func GetStudyProgram(code string) string {
	var statusGetStudyProgram string
	var studyPrograms []string

	studyPrograms = strings.Split(StudentStudyPrograms, ", ")

	if code == "" {
		statusGetStudyProgram = "Code is undefined!"
	} else {
		for i := 0; i < len(studyPrograms); i++ {
			if strings.Split(studyPrograms[i], "_")[0] == code {
				statusGetStudyProgram = strings.Split(studyPrograms[i], "_")[1]
				break
			} else {
				statusGetStudyProgram = "Study tidak ditemukan"
			}
		}
	}
	return statusGetStudyProgram // TODO: replace this
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
