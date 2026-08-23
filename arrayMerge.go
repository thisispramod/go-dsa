package main

// arr1 : 1,2,3,4,5 arr2 : 6,7,8,9,10

func ArrayMerge(arr1 []int, arr2 []int) []int {
	for i := 0; i < len(arr2); i++ {
		arr1 = append(arr1, arr2[i])
	}
	return arr1
}

func SortedArrMerge(arr1 []int, arr2 []int) []int {

	n, m := len(arr1), len(arr2)

	i := 0
	j := 0
	arr := make([]int, 0, n+m)
	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	if i < n {
		arr = append(arr, arr1[i:]...)

	}
	if j < m {
		arr = append(arr, arr2[j:]...)
	}
	return arr
}
