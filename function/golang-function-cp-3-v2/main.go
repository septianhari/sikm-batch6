package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	// Split string berdasarkan spasi, koma, dan titik koma
	nameList := strings.FieldsFunc(names, func(r rune) bool {
		return r == ' ' || r == ',' || r == ';'
	})

	// Inisialisasi nama terpendek dan panjangnya
	shortestName := ""
	shortestLength := 0

	for _, name := range nameList {
		// Hilangkan whitespace di awal dan akhir nama
		name = strings.TrimSpace(name)
		// Panjang nama
		nameLength := len(name)

		if shortestName == "" || nameLength < shortestLength {
			shortestName = name
			shortestLength = nameLength
		} else if nameLength == shortestLength && name < shortestName {
			// Jika panjang nama sama, namun nama baru lebih kecil secara alfabetis
			shortestName = name
		}
	}

	return shortestName
}

func main() {
	// Test case 1
	names1 := "Hanif Joko Tio Andi Budi Caca Hamdan"
	fmt.Println(FindShortestName(names1)) // Output: "Tio"

	// Test case 2
	names2 := "Budi;Tia;Tio"
	fmt.Println(FindShortestName(names2)) // Output: "Tia"
}
