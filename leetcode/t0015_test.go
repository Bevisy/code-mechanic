package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			[]int{-1, 0, 1},
			[]int{-1, -1, 2},
		}},
	}

	for _, test := range tests {
		got := threeSum(test.input)
		fmt.Printf("got: %v", got)
		if len(got) == len(test.input) {
			for i, v := range test.expected {
				for j := range v {
					if got[i][j] != test.expected[i][j] {
						t.Errorf("input: %v, expected: %v, but got: %d", test.input, test.expected, got)
					}
				}
			}
		} else {
			t.Errorf("got is not the same as expected.")
		}
	}
}
