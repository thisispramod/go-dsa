package main

import "fmt"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)
	arrm := make([]int, 0, m+n)
	i := 0
	j := 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			arrm = append(arrm, nums1[i])
			i++
		} else {
			arrm = append(arrm, nums2[j])
			j++
		}
	}

	if i < m {
		arrm = append(arrm, nums1[i:]...)
	}
	if j < n {
		arrm = append(arrm, nums2[j:]...)
	}

	numlength := len(arrm)
	fmt.Println(arrm)
	if numlength%2 == 0 {
		first := arrm[numlength/2]
		second := arrm[(numlength/2 - 1)]
		return float64(first+second) / 2
	} else {
		return float64(arrm[numlength] / 2)
	}

}
