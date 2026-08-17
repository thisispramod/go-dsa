package main

// pramod //domarp
/*
func ReverseString(str string) string {
	var newstr string
	for i := len(str) - 1; i >= 0; i-- {
		newstr += string(str[i])
	}
	return newstr
}
*/

func ReversStringPointer(str string) string {

	i := 0
	j := len(str) - 1

	b := []byte(str)
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}
