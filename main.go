package main

import (
	// "fmt"
	"fmt"
	"learn-golang/practice"
)

func main() {
	// fmt.Printf("This is the main file")
	// var sum = practice.SumOfIntegers(1, 2, 3, 4)
	// fmt.Printf("sum: %d", sum)

	name := "Ana"
	isPalindrome := practice.CheckIfPalindrome(name)

	if isPalindrome {
		fmt.Printf("%s is  palindrome", name)
	} else {
		fmt.Printf("%s is not a palindrome", name)
	}
}
