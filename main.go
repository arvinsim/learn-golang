package main

import (
	"fmt"
	"learn-golang/leetcode"
	"learn-golang/practice"
)

func main() {
	//outputResult("saippuakivikauppias")
	//outputResult("foobar")

	nums := []int{1, 7, 3, 6, 5, 6}
	result := leetcode.PivotIndex(nums)

	fmt.Printf("%d", result)
}

func outputResult(text string) {
	isPalindrome := practice.CheckIfPalindrome(text)

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome\n", text)
	} else {
		fmt.Printf("'%s' is a not palindrome\n", text)
	}
}
