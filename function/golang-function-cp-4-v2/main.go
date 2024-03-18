package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk mencari data yang mirip dengan input
func FindSimilarData(input string, data ...string) string {
	var similarData []string

	// Iterasi melalui semua data yang diberikan
	for _, d := range data {
		// Jika data mengandung input, tambahkan ke dalam similarData
		if strings.Contains(d, input) {
			similarData = append(similarData, d)
		}
	}

	// Gabungkan semua data yang mirip dengan input menjadi satu string dengan dipisahkan koma
	return strings.Join(similarData, ",")
}

func main() {
	// Test case 1
	input1 := "mobil"
	data1 := []string{"mobil APV", "mobil Avanza", "motor matic", "motor gede"}
	fmt.Println("Test Case 1:")
	fmt.Println("Input:", input1, "Data:", data1)
	fmt.Println("Output:", FindSimilarData(input1, data1...))

	// Test case 2
	input2 := "motor"
	data2 := []string{"mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"}
	fmt.Println("\nTest Case 2:")
	fmt.Println("Input:", input2, "Data:", data2)
	fmt.Println("Output:", FindSimilarData(input2, data2...))
}
