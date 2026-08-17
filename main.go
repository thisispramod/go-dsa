package main

import "fmt"

func main() {

	/*
		resA, resB := SwapXor(10, 20)
		fmt.Println("a swap to b:", resA, resB)
	*/

	/*
		fmt.Println("count of string is ", UniqueFirst("swiss"))
	*/

	/*
		arr := []int{2, 7, 11, 15}
		target := 9
		fmt.Println("Two sum problem indices", TwoSum(arr, target))
	*/

	/*	Find Factorial By Taking Input From Console eg: 5, Factorial is 120 (5*4*3*2*1)*/
	/*
		var userInput int
		fmt.Print("Enter Your value: ")
		fmt.Scan(&userInput)
		factorialValue, err := FactorialFind(userInput)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Entered Value is %d and its Factorial is %d \n", userInput, factorialValue)
	*/
	/*

		Input: "()[]{}"
		Output: true

		Input: "(]"
		Output: false

		Input: "([)]"
		Output: false  (order galat hai — '(' ke andar '[' hai, lekin ']' pehle aa gaya)

		Input: "{[]}"
		Output: true

	*/
	/*
		var userString string
		fmt.Print("Enter Your string: ")
		fmt.Scan(&userString)
		fmt.Println("Test Case:", Matching(userString))
	*/
	/*
		fmt.Println("substring count", LongestSubstring("pwwkewe"))
	*/

	/* fmt.Println("string return ", ReverseString("pramod")) */
	fmt.Println("using two pointer", ReversStringPointer("root"))

}
