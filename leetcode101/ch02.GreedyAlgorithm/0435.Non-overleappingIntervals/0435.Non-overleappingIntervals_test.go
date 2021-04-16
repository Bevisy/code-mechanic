package leetcode

import (
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
		{[][]int{{1, 4}, {2, 3}, {1, 2}, {1, 3}}, 2},
	}

	for _, test := range tests {
		got := eraseOverlapIntervals(test.input)
		if test.expect != got {
			t.Errorf("input: %v, expect: %d, but got: %d.", test.input, test.expect, got)
		}
	}
}
