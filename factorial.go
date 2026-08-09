package main

import "fmt"

func FactorialFind(num int) (int, error) {
	if num < 0 {
		return 0, fmt.Errorf("factorial is not defined for negative numbers")
	}
	factorial := 1
	for i := 2; i <= num; i++ {
		factorial *= i
	}
	return factorial, nil
}
