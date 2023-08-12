package leetcode

import "testing"

func TestNumberOfSteps(t *testing.T) {
	test_cases := []struct {
		input    int
		expected int
	}{
		{input: 14, expected: 6},
		{input: 8, expected: 4},
	}

	for _, test_case := range test_cases {
		result := NumberOfSteps(test_case.input)

		if result != test_case.expected {
			t.Errorf("Expected NumbrOfSteps(%v) to return %d", test_case.input, test_case.expected)
		}
	}
}
