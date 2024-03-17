package main

import "fmt"

func GraduateStudent(score, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}

func main() {
	// Test Case 1
	score1, absent1 := 100, 4
	fmt.Printf("Test Case 1: score = %d, absent = %d\n", score1, absent1)
	fmt.Println("Expected Output:", GraduateStudent(score1, absent1))

	// Test Case 2
	score2, absent2 := 80, 5
	fmt.Printf("Test Case 2: score = %d, absent = %d\n", score2, absent2)
	fmt.Println("Expected Output:", GraduateStudent(score2, absent2))
}
