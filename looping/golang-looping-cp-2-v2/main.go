package main

import "fmt"

func ReverseString(str string) string {

	var reverse string

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}

	//fmt.Println(reverse)

	var final string

	for i := 0; i < len(reverse); i++ {
		if (i+1 < len(reverse) && reverse[i+1] == ' ') || reverse[i] == ' ' {
			final += string(reverse[i])
		} else {
			final += string(reverse[i]) + "_"
		}
	}

	return final[:len(final)-1]
}

//gunakan untuk melakukan debug

func main() {
	fmt.Println(ReverseString("Hello World"))
	fmt.Println(ReverseString("I am a student"))
}

//"I am a student"
