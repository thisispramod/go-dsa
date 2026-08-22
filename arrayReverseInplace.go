package main

func ArrayReverseInplace(arr []int) []int {

	i := 0
	j := len(arr) - 1

	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i++
		j--
	}

	return arr

}
