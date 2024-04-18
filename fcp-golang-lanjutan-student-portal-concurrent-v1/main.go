package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) (string, error)
	Register(id string, name string, studyProgram string) (string, error)
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) (string, error)
	ImportStudents(filenames []string) error
	SubmitAssignments(numAssignments int)
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	failedLoginAttempts  map[string]int // map untuk melacak percobaan login yang gagal
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		failedLoginAttempts: make(map[string]int), // inisialisasi map failedLoginAttempts
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	sm.Lock()
	defer sm.Unlock()

	// Periksa apakah ID telah terblokir
	if attempts, ok := sm.failedLoginAttempts[id]; ok && attempts >= 3 {
		return "", fmt.Errorf("Login gagal: Batas maksimum login terlampaui")
	}

	// Temukan siswa dengan ID yang sesuai
	for _, student := range sm.students {
		if student.ID == id && student.Name == name {
			// Reset percobaan login gagal jika login berhasil
			delete(sm.failedLoginAttempts, id)
			return fmt.Sprintf("Selamat datang %s! Kamu terdaftar di program studi: %s.", name, sm.studentStudyPrograms[student.StudyProgram]), nil
		}
	}

	// Jika tidak ada siswa yang sesuai, tambahkan percobaan login gagal
	sm.failedLoginAttempts[id]++
	return "", fmt.Errorf("Login gagal: Data mahasiswa tidak ditemukan")
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// Proses pendaftaran...
	return "Registrasi berhasil", nil
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	// Dapatkan nama program studi berdasarkan kode
	if program, ok := sm.studentStudyPrograms[code]; ok {
		return program, nil
	}
	return "", fmt.Errorf("Program studi tidak ditemukan")
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	// Ubah data mahasiswa...
	return "Perubahan berhasil", nil
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		// Ubah program studi mahasiswa...
		return nil
	}
}

func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
	// Import data mahasiswa dari file CSV...
	return nil
}

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// Proses pengiriman tugas yang memakan waktu...
	time.Sleep(3000 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {
	fmt.Println("=== Submit Assignment ===")
	fmt.Printf("Enter the number of assignments you want to submit: %d\n", numAssignments)

	// Simulasi pengiriman tugas
	workCh := make(chan int)
	for i := 0; i < 3; i++ {
		go func(workerID int) {
			for assignment := range workCh {
				fmt.Printf("Worker %d: Processing assignment %d\n", workerID, assignment)
				sm.SubmitAssignmentLongProcess()
				fmt.Printf("Worker %d: Finished assignment %d\n", workerID, assignment)
			}
		}(i + 1)
	}

	for i := 1; i <= numAssignments; i++ {
		workCh <- i
	}
	close(workCh)

	fmt.Printf("Submitting %d assignments...\n", numAssignments)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println(msg)
			}
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			// Handle registration...
		case "3":
			// Handle getting study program...
		case "4":
			// Handle modifying student...
		case "5":
			// Handle bulk import...
		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignmentsInput, _ := reader.ReadString('\n')
			numAssignmentsInput = strings.TrimSpace(numAssignmentsInput)
			numAssignments, err := strconv.Atoi(numAssignmentsInput)
			if err != nil {
				fmt.Println("Error: Please enter a valid number")
				break
			}
			manager.SubmitAssignments(numAssignments)
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}
	}
}
