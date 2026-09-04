package main

func StringConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var ans string
	n := len(s)
	charinsection := 2 * (numRows - 1)
	for i := 0; i < numRows; i++ {
		var index int = i
		for index < n {
			ans += string(s[index])
			if i != 0 && i != numRows-1 {
				charinbetween := charinsection - 2*i
				secondindex := index + charinbetween
				if secondindex < n {
					ans += string(s[secondindex])
				}
			}
			index += charinsection

		}
	}
	return ans
}
