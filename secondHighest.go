package main

import "math"

func SecondHighest(arr []int) int {

	if len(arr) < 2 {
		return math.MinInt
	}

	firstMax := math.MinInt
	secondMax := math.MinInt
	// 1,2,3,4,5
	for i := 0; i < len(arr); i++ {
		if firstMax < arr[i] {
			secondMax = firstMax
			firstMax = arr[i]
		} else if secondMax != firstMax && firstMax > secondMax {
			secondMax = arr[i]
		}
	}

	return secondMax
}
