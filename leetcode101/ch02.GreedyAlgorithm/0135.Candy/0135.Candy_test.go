package leetcode

import "testing"

func TestCandy(t *testing.T) {
	var tests = []struct {
		test   []int
		expect int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13},
	}

	for _, test := range tests {
		got := candy(test.test)
		if got != test.expect {
			t.Errorf("input: %v, expect: %d, but got: %d\n", test.test, test.expect, got)
		}
	}
}
