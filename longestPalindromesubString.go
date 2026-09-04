/*
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
forgeeksskeegfor

	└──────────┘
	 geeksskeeg
*/
package main

// forgeeksskeegfor
func LogestPalindromesubString(str string) string {

	if len(str) == 0 {
		return ""
	}

	longestStart, longestEnd := 0, 0

	expand := func(left, right int) {
		for left >= 0 && right < len(str) && str[left] == str[right] {
			if (right - left) > (longestEnd - longestStart) {
				longestStart = left
				longestEnd = right
			}
			left--
			right++
		}
	}

	for i := 0; i < len(str); i++ {
		expand(i, i)
		expand(i, i+1)
	}
	return str[longestStart : longestEnd+1]
}
