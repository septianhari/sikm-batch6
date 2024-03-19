package main

func MapToSlice(mapData map[string]string) [][]string {
	result := [][]string{}

	for key, value := range mapData {
		result = append(result, []string{key, value})
	}

	return result
}

func main() {
	// Contoh penggunaan fungsi MapToSlice
	mapData := map[string]string{"hello": "world", "John": "Doe", "age": "14"}
	output := MapToSlice(mapData)
	println(output)
}
