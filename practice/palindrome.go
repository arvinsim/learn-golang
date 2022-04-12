package main

import (
	"fmt"
	"strings"
)

func main() {
	isPalindrome := false
	foo := "saippuakivikauppias"
	length := len(foo)
	reverse := []string{}

	for i := length - 1; i >= 0; i-- {
		reverse = append(reverse, string(foo[i]))
	}

	if foo == strings.Join(reverse, "") {
		isPalindrome = true
	}

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome", foo)
	} else {
		fmt.Printf("'%s' is a not palindrome", foo)
	}
}
