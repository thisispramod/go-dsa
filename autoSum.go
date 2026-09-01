package main

func SumOfGiven(num uint64) uint64 {
	/*
		var sum int
				for i := 0; i <= num; i++ {
					sum += i
				}
			return sum
	*/
	return num * (num + 1) / 2
}
