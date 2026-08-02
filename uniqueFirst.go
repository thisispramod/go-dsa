/* Input: "swiss"
Output: 'w'  (kyunki 's' aur 'i' repeat hote hain, 'w' pehla unique hai)

Input: "aabbcc"
Output: -1 (koi bhi unique nahi hai) */

package main

func UniqueFirst(str string) int {
	counts := make(map[rune]int)

	for _, char := range str {
		counts[char]++
	}

	for i, char := range str {
		if counts[char] == 1 {
			return i
		}
	}

	return -1
}
