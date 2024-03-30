package main

import (
	"fmt"
	"strings"
)

// Package internal berisi struct Calculator yang dapat menjalankan 4 operasi sederhana
// yaitu Add, Subtract, Multiply, Divide. Dan dapat memanggil fungsi Result untuk mendapatkan hasil dari operasi matematika yang telah dilakukan.
//package internal

type Calculator struct {
	total float32
}

// Method untuk menambahkan angka ke dalam total
func (c *Calculator) Add(num float32) {
	c.total += num
}

// Method untuk mengurangkan angka dari total
func (c *Calculator) Subtract(num float32) {
	c.total -= num
}

// Method untuk mengalikan angka dengan total
func (c *Calculator) Multiply(num float32) {
	c.total *= num
}

// Method untuk membagi total dengan angka
func (c *Calculator) Divide(num float32) {
	if num != 0 {
		c.total /= num
	}
}

// Method untuk mendapatkan hasil dari operasi matematika yang telah dilakukan
func (c *Calculator) Result() float32 {
	return c.total
}

func AdvanceCalculator(calculate string) float32 {
	// Pisahkan string menjadi token-token (angka dan operator) menggunakan spasi sebagai delimiter
	tokens := strings.Fields(calculate)

	// Inisialisasi kalkulator
	calculator := &Calculator{}

	// Jika tidak ada token, kembalikan 0.0
	if len(tokens) == 0 {
		return calculator.Result()
	}

	// Inisialisasi nilai awal dengan token pertama
	num, _ := parseNumber(tokens[0])
	calculator.Add(num)

	// Loop melalui token-token yang tersisa untuk melakukan operasi matematika
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		num, _ := parseNumber(tokens[i+1])

		switch operator {
		case "+":
			calculator.Add(num)
		case "-":
			calculator.Subtract(num)
		case "*":
			calculator.Multiply(num)
		case "/":
			calculator.Divide(num)
		}
	}

	// Kembalikan hasil dari kalkulasi
	return calculator.Result()
}

// Fungsi untuk mengubah string menjadi float32
func parseNumber(str string) (float32, error) {
	var num float32
	_, err := fmt.Sscanf(str, "%f", &num)
	return num, err
}

func main() {
	fmt.Println(AdvanceCalculator("3 * 4 / 2 + 10 - 5"))                    // Output: 11.0
	fmt.Println(AdvanceCalculator("10 / 4 + 100"))                          // Output: 102.5
	fmt.Println(AdvanceCalculator("10 + 10 + 10 + 10 + 12 + 12 + 12 + 12")) // Output: 72.0
	fmt.Println(AdvanceCalculator("10"))                                    // Output: 10.0
	fmt.Println(AdvanceCalculator(""))                                      // Output: 0.0
}
