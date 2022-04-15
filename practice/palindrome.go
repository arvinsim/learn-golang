package practice

import (
	"fmt"
	"strings"
)

func main() {
	outputResult("saippuakivikauppias")
	outputResult("foobar")
}

func CheckIfPalindrome(text string) bool {
	length := len(text)
	reverse := []string{}

	for i := length - 1; i >= 0; i-- {
		reverse = append(reverse, string(text[i]))
	}

	if text == strings.Join(reverse, "") {
		return true
	}
	return false
}

func outputResult(text string) {
	isPalindrome := CheckIfPalindrome(text)

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome\n", text)
	} else {
		fmt.Printf("'%s' is a not palindrome\n", text)
	}
}
