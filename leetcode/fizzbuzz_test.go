package leetcode

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	test_cases := []struct {
		expected []string
		n        int
	}{
		{expected: []string{"1", "2", "Fizz"}, n: 3},
		{expected: []string{"1", "2", "Fizz", "4", "Buzz"}, n: 5},
		{expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, n: 15},
	}

	for _, test_case := range test_cases {
		result := FizzBuzz(test_case.n)

		if !reflect.DeepEqual(result, test_case.expected) {
			t.Errorf("Expected FizzBuzz(%v) => (%v) to return %v", test_case.n, result, test_case.expected)
		}
	}
}
