package main

/*
Input: s = "pwwkew"
Output: 3
Explanation: Longest unique substring hai "wke", length 3.
*/

func LongestSubstring(s string) int {

	start := 0 // current window ka start pointer
	maxLength := 0
	lastSeen := map[byte]int{} // pwwkew
	for i := 0; i < len(s); i++ {
		char := s[i]
		_, ok := lastSeen[char]
		if ok {
			if lastSeen[char] >= start {
				start = lastSeen[char] + 1
			}
		}

		lastSeen[char] = i

		// Current window ki length nikaalo
		currentLength := i - start + 1 // 1, 2,2

		// Agar ye ab tak ka best hai, to update karo
		if currentLength > maxLength {
			maxLength = currentLength // 1,2
		}
	}
	return maxLength
}
