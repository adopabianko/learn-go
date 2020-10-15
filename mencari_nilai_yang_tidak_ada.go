package main

import (
	"fmt"
	"sort"
)

func main() {
	var n = []int{3, 5, 2, 1, 9}

	min := MinSlice(n)
	max := MaxSlice(n)

	var outputSlice []int

	for i := min; i <= max; i++ {
		found := FindSlice(n, i)
		if !found {
			outputSlice = append(outputSlice, i)
		}
	}

	fmt.Println(outputSlice)
}

func MinSlice(v []int) int {
	sort.Ints(v)

	return v[0]
}

func MaxSlice(v []int) int {
	sort.Ints(v)

	return v[len(v)-1]
}

func FindSlice(slice []int, val int) bool {
	for item := range slice {
		if item == val {
			return true
		}
	}

	return false
}
