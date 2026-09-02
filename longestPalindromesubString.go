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

func LogestPalindromesubString(str string) bool {
	var status bool = true
	i := 0
	j := len(str) - 1
	var strlen int
	for i < j {
		if str[i] != str[j] {
			status = false
			strlen = 1
			// i++
			// j--
		} else if str[i] == str[j] {
			// status = true
			if strlen >= 2 {
				status = true
			}
			strlen++
		}
		i++
		j--
	}
	return status
}
