package main

import (
	"strconv"
	"strings"
)

func FabonacciSeries(n int) string {
	faboSlice := make([]string, n)
	faboSlice[0] = "0"
	faboSlice[1] = "1"
	first := 0
	second := 1
	for i := 2; i < n; i++ {
		first, second = second, first+second
		faboSlice[i] = strconv.Itoa(second)
	}
	str := strings.Join(faboSlice, ",")
	return str

}
