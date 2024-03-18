package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowels := "aiueo"
	vowelCount, consonantCount := 0, 0

	for _, char := range str {
		if strings.Contains(vowels, strings.ToLower(string(char))) {
			vowelCount++
		} else if char != ' ' {
			consonantCount++
		}
	}

	return vowelCount, consonantCount, vowelCount == 0 || consonantCount == 0
}

func main() {
	// Test cases
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"kopi", "2, 2, false"},
		{"bbbbb ccccc", "0, 10, true"},
		{"Hidup Itu Indah", "6, 7, false"},
	}

	for _, tc := range testCases {
		vowels, consonants, noVowelOrConsonant := CountVowelConsonant(tc.input)
		output := fmt.Sprintf("%d, %d, %t", vowels, consonants, noVowelOrConsonant)
		if output == tc.expectedOutput {
			fmt.Printf("Input: %q\nOutput: %s\nExpected: %s\nResult: Passed\n\n", tc.input, output, tc.expectedOutput)
		} else {
			fmt.Printf("Input: %q\nOutput: %s\nExpected: %s\nResult: Failed\n\n", tc.input, output, tc.expectedOutput)
		}
	}
}
