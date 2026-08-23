package main

// str = racecar , taxxxat
// output = true
func checkPalindromestring(s string) bool {
	i := 0
	str := []rune(s)
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true

}
