package practice

import (
	"testing"
)

func TestCheckIfPalindrome(t *testing.T) {
	tables := []struct {
		isPalindrome bool
		text         string
	}{
		{isPalindrome: true, text: "asa"},
		{isPalindrome: false, text: "foobar"},
	}

	for _, item := range tables {
		result := CheckIfPalindrome(item.text)

		if result != item.isPalindrome {
			t.Errorf("Expected CheckIfPalindrome(%s) to return %t", item.text, item.isPalindrome)
		}
	}
}
