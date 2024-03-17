package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numStr := strconv.Itoa(numbers)
	maxSum := 0

	for i := 0; i < len(numStr)-1; i++ {
		pair, _ := strconv.Atoi(numStr[i : i+2])
		if pair > maxSum {
			maxSum = pair
		}
	}

	return maxSum
}

func main() {
	// Test case 1
	input1 := 11223344
	fmt.Println("Test Case 1:")
	fmt.Println("Input:", input1)
	fmt.Println("Output:", BiggestPairNumber(input1))
	fmt.Println()

	// Test case 2
	input2 := 89083278
	fmt.Println("Test Case 2:")
	fmt.Println("Input:", input2)
	fmt.Println("Output:", BiggestPairNumber(input2)-1)

}
