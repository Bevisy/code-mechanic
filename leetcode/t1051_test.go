package leetcode

import "testing"

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
	}

	for _, test := range tests {
		got := heightChecker(test.input)
		if got != test.expect {
			t.Errorf("Input: %v, expect: %d, Got: %d\n", test.input,
				test.expect, got)
		}
	}
}
