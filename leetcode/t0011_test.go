package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, test := range tests {
		got := maxArea(test.input)
		if got != test.expected {
			t.Errorf("input: %v, expected: %d, but got: %d", test.input, test.expected, got)
		}
	}
}
