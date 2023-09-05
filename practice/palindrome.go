package practice

import (
	"strings"
)

func CheckIfPalindrome(text string) bool {
	length := len(text)
	reverse := []string{}
	lowercaseText := strings.ToLower(text)

	for i := length - 1; i >= 0; i-- {
		character := string(lowercaseText[i])
		reverse = append(reverse, character)
	}

	return lowercaseText == strings.Join(reverse, "")
}
