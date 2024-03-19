package main

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{} // jika data kosong, kembalikan slice kosong
	}

	totalProfit := make([]int, len(data[0]))

	for _, branchData := range data {
		for i, monthData := range branchData {
			profit := monthData[0] - monthData[1]
			totalProfit[i] += profit
		}
	}

	return totalProfit
}
