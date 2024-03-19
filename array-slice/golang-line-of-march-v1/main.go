package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	fmt.Println("Before", height)
	sort.Slice(height, func(i, j int) bool {
		return height[j] > height[i]
	})
	return height
}
