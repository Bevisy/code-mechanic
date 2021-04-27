package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, test := range tests {
		got := twoSum(test.input, test.target)
		for i := range got {
			if got[i] != test.expect[i] {
				t.Errorf("input: %v, want: %v, but got: %v\n", test.input,
					test.expect, got)
			}
		}
	}
}
