package main

// [0,0,1,0,1,1,1,0,0,0,1,1,0,1,1,0,1]
func MoveallzerosLeft(arr []int) []int {
	i := 0
	j := len(arr) - 1

	for i < j {
		if arr[i] == 0 {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	return arr
}
