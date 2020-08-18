package leetcode

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range tests {
		got := twoSum(test.input, 9)
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("input: %v, expected: %v, but got: %v", test.input, test.expected, got)
			}
		}
	}
}
