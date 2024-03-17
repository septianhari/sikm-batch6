package main

import "fmt"

// GetPredicate mengembalikan predikat berdasarkan nilai rata-rata dari 4 mata pelajaran
func GetPredicate(math, science, english, indonesia int) string {
	average := (math + science + english + indonesia) / 4

	switch {
	case average == 100:
		return "Sempurna"
	case average >= 90:
		return "Sangat Baik"
	case average >= 80:
		return "Baik"
	case average >= 70:
		return "Cukup"
	case average >= 60:
		return "Kurang"
	case average <= 60:
		return "Sangat kurang"
	default:
		return "Bahlul"
	}
}

func main() {
	// Test case 1
	predikat1 := GetPredicate(50, 80, 100, 60)
	fmt.Println("Test Case 1:", predikat1) // Output: Cukup

	// Test case 2
	predikat2 := GetPredicate(100, 100, 100, 100)
	fmt.Println("Test Case 2:", predikat2) // Output: Sempurna
}
