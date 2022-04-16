package main

import (
	"fmt"
	"learn-golang/practice"
)

func main() {
	outputResult("saippuakivikauppias")
	outputResult("foobar")
}

func outputResult(text string) {
	isPalindrome := practice.CheckIfPalindrome(text)

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome\n", text)
	} else {
		fmt.Printf("'%s' is a not palindrome\n", text)
	}
}
