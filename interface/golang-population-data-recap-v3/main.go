package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	result := []map[string]any{}
	for _, v := range data {
		curResult := make(map[string]any)
		tokens := strings.Split(v, ";")
		name := tokens[0]
		age, _ := strconv.Atoi(tokens[1])
		address := tokens[2]
		height, _ := strconv.ParseFloat(tokens[3], 64)
		isMarried, _ := strconv.ParseBool(tokens[4])

		curResult["name"] = name
		curResult["age"] = age
		curResult["address"] = address
		if tokens[3] != "" {
			curResult["height"] = height
		}
		if tokens[4] != "" {
			curResult["isMarried"] = isMarried
		}

		result = append(result, curResult)
	}
	return result // TODO: replace this
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	output := PopulationData(data)
	fmt.Println(output)
}
