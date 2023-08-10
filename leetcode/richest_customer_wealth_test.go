package leetcode

import "testing"

func TestGetRichestCustomerWealth(t *testing.T) {
	tables := []struct {
		expected int
		matrix   [][]int
	}{
		{expected: 6, matrix: [][]int{{1, 2, 3}, {3, 2, 1}, {1, 1, 1}}},
		{expected: 10, matrix: [][]int{{1, 5}, {7, 3}, {3, 5}}},
		{expected: 0, matrix: [][]int{{0}}},
	}

	for _, item := range tables {
		result := GetRichestCustomerWealth(item.matrix)

		if result != item.expected {
			t.Errorf("Expected GetRichestCustomerWealth(%v) to return %d", item.matrix, item.expected)
		}
	}
}
