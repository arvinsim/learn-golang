package leetcode

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	test_cases := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{ransomNote: "a", magazine: "b", expected: false},
		{ransomNote: "aa", magazine: "ab", expected: false},
		{ransomNote: "aa", magazine: "aab", expected: true},
	}

	for _, test_case := range test_cases {
		result := CanConstruct(test_case.ransomNote, test_case.magazine)

		if result != test_case.expected {
			t.Errorf("Expected CanConstruct(%s, %s) to return %v", test_case.ransomNote, test_case.magazine, test_case.expected)
		}
	}
}
