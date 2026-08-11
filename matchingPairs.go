package main

func Matching(str string) bool {

	matchingpairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < len(str); i++ {

		char := str[i]

		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}

			topElement := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			if topElement != matchingpairs[char] {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
