package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var vocalCount, consonantCount int
	var isValid bool

	for _, c := range str {
		cc := strings.ToLower(string(c))

		if cc == "a" || cc == "i" || cc == "u" || cc == "e" || cc == "o" {
			vocalCount += 1
		} else if cc >= "a" && cc <= "z" {
			consonantCount += 1
		}
	}

	if vocalCount == 0 || consonantCount == 0 {
		isValid = true
	} else {
		isValid = false
	}

	return vocalCount, consonantCount, isValid // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("bbbb cccc"))
	fmt.Println(CountVowelConsonant("aaa iii"))

}
