/*
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]  (kyunki nums[0] + nums[1] = 2 + 7 = 9)
*/
package main

func TwoSum(arrs []int, target int) []int {
	lookup := make(map[int]int)

	for currentIndex, num := range arrs {
		diff := target - num
		value, ok := lookup[diff]
		if ok {
			return []int{value, currentIndex}
		}
		lookup[num] = currentIndex
	}
	return nil

}
