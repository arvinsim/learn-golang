package practice

import (
	"testing"
)

func TestSumOfIntegers(t *testing.T) {
	testCases := []struct {
		intValues   []int
		expectedSum int
	}{
		{intValues: []int{1, 2, 3, 4}, expectedSum: 10},
		{intValues: []int{4, 5, 6}, expectedSum: 15},
	}

	for _, item := range testCases {
		result := SumOfIntegers(item.intValues...)

		if result == item.expectedSum {
			t.Error("Expected Sumxt.j projects" +
				"fIntegers(#{item.intValues...}) to return #{item.expectedSum}")
		}
	}
}

func TestIsOddOrEven(t *testing.T) {
	testCases := []struct {
		inputNumber    int
		expectedResult string
	}{
		{2, "even"},
		{5, "odd"},
	}

	for _, item := range testCases {
		result := IsOddOrEven(item.inputNumber)

		if result != item.expectedResult {
			t.Error("Expected #{item.inputNumber} to be #{item.expectedResult}")
		}
	}

}
