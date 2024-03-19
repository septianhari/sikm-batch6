package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	DataSplit := make([][]string, len(data))
	for i := 0; i < len(data); i++ {
		DataSplit[i] = append(DataSplit[i], strings.Split(data[i], "-")...)
	}

	// untuk mendapatkan key untuk map
	KeyDict := []string{}
	rev := ""
	for _, key := range DataSplit[:] {
		if rev != key[0] {
			rev = key[0]
			KeyDict = append(KeyDict, key[0])
		}
	}

	// membuat map dengan panjang sesuai keydict yaitu disini ada 3
	result := make(map[string][]string, len(KeyDict))

	for key := 0; key < len(KeyDict); key++ {
		for i := 0; i < len(DataSplit); i++ {
			// j ini adalah data berikutnya atau "last"
			for j := i + 1; j < len(DataSplit); j++ {
				if DataSplit[i][0] == KeyDict[key] && DataSplit[j][0] == KeyDict[key] {
					if j%2 == 0 {
						break
					}

					if DataSplit[i][0] == "phone" && DataSplit[j][0] == "phone" {
						result[KeyDict[key]] = append(result[KeyDict[key]], DataSplit[i][3])
						result[KeyDict[key]] = append(result[KeyDict[key]], DataSplit[j][3])
						break
					}

					// menggabungkan string berdasarkan index yang sama (account dan address)
					itemString := DataSplit[i][3] + " " + DataSplit[j][3]

					n := 0

					for _, SameName := range result[KeyDict[key]] {
						Check := false
						Check = strings.Contains(SameName, itemString)

						if Check {
							n++
							break
						}
					}

					if n == 1 {
						break
					}
					result[KeyDict[key]] = append(result[KeyDict[key]], itemString)
					break
				}
			}
		}
	}
	return result // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
