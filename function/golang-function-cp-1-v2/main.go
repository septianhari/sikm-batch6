package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	// Membuat map untuk mengonversi bulan dari angka menjadi nama bulan dalam bahasa Inggris
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	// Mengonversi day menjadi string dengan format dua digit
	dayStr := fmt.Sprintf("%02d", day)

	// Mengambil nama bulan dari map months
	monthStr := months[month]

	// Mengonversi year menjadi string
	yearStr := fmt.Sprintf("%d", year)

	// Menggabungkan day, month, dan year dalam format yang diinginkan
	return fmt.Sprintf("%s-%s-%s", dayStr, monthStr, yearStr)
}

func main() {
	// Contoh penggunaan fungsi DateFormat
	fmt.Println(DateFormat(1, 1, 2020))   // Output: 01-January-2020
	fmt.Println(DateFormat(31, 12, 2020)) // Output: 31-December-2020
	fmt.Println(DateFormat(6, 5, 2021))   // Output: 06-May-2020
}
