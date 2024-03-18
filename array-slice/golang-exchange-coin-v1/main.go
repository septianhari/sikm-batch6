package main

import "fmt"

func ExchangeCoin(amount int) []int {
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0)

	for _, coin := range coins {
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}

	return result
}

func main() {
	fmt.Println(ExchangeCoin(1752)) // Output: [1000 500 200 50 1 1]
	fmt.Println(ExchangeCoin(5000)) // Output: [1000 1000 1000 1000 1000]
	fmt.Println(ExchangeCoin(1234)) // Output: [1000 200 20 10 1 1 1 1]
	fmt.Println(ExchangeCoin(0))    // Output: []
}
