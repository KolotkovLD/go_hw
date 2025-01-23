package main

import "fmt"

func getReversed(num int) int {
	remains := num
	reversed := 0
	multiplier := 1
	for remains > 0 {
		tail := remains % 10
		reversed = reversed*10 + tail
		remains = (remains - tail) / 10
		multiplier = multiplier * 10
	}
	return reversed
}

func isPalindrome(firstNum, secondNum int) bool {
	secondReversed := getReversed(secondNum)
	if firstNum == secondReversed {
		return true
	}
	return false
}

func main() {
	//fmt.Println(getReversed(1543))
	fmt.Printf("%t", isPalindrome(1033, 331))
}
