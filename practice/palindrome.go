package practice

import (
	"strings"
)

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
