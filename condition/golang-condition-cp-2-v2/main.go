package main

import (
	"fmt"
)

func BMICalculator(gender string, height int) float64 {
	var bmi float64
	if gender == "laki-laki" {
		bmi = float64(height-100) - float64((height-100)*10)/100
	} else if gender == "perempuan" {
		bmi = float64(height-100) - float64((height-100)*15)/100
	}
	return bmi
} // TODO: replace this

// gunakan untuk melakukan debug
func main() {
	// Test Case 1
	gender1 := "laki-laki"
	height1 := 170
	fmt.Printf("Test Case 1:\nInput: gender = %s, height = %d\nOutput: %v\n\n", gender1, height1, BMICalculator(gender1, height1))

	// Test Case 2
	gender2 := "perempuan"
	height2 := 165
	fmt.Printf("Test Case 2:\nInput: gender = %s, height = %d\nOutput: %v\n\n", gender2, height2, BMICalculator(gender2, height2))
}
