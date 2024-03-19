package main

func SchedulableDays(villager [][]int) []int {
	availableDate := make(map[int]int)
	for _, v := range villager {
		for _, tanggal := range v {

			availableDate[tanggal] = availableDate[tanggal] + 1

		}
	}

	result := []int{}
	for k, v := range availableDate {
		if v == len(villager) {
			result = append(result, k)
		}
	}

	return result
}
