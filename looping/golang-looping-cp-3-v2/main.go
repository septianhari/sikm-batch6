package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z

	var countingCadel = 0

	for i := 0; i < len(text); i++ {
		if text[i] == 'R' ||
			text[i] == 'S' ||
			text[i] == 'T' ||
			text[i] == 'Z' ||
			text[i] == 'r' ||
			text[i] == 's' ||
			text[i] == 't' ||
			text[i] == 'z' {
			countingCadel++
		}
	}
	return countingCadel // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
	fmt.Println(CountingLetter("Zebra Zig Zag"))
	fmt.Println(CountingLetter("Ular Melingkar Di atas Pagarrrrr! roar"))
}
