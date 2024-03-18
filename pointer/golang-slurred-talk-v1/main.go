package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	// Mapping huruf yang akan diganti menjadi huruf L
	replacements := strings.NewReplacer("s", "l", "S", "L", "r", "l", "R", "L", "z", "l", "Z", "L")
	// Lakukan penggantian huruf sesuai dengan mapping
	*words = replacements.Replace(*words)
}

func main() {
	// Test case
	testCases := []struct {
		words   string
		expects string
	}{
		{"Steven", "Lteven"},
		{"Saya Steven", "Laya Lteven"},
		{"Saya Steven, saya suka menggoreng telur dan suka hewan zebra", "Laya Lteven, laya luka menggoreng telur dan luka hewan lebra"},
		{"", ""},
	}

	// Uji setiap kasus uji dan cetak hasilnya
	for _, tc := range testCases {
		words := tc.words
		SlurredTalk(&words)
		fmt.Printf("Input: %s, \n Output: %s\n", tc.words, words)
	}
}
