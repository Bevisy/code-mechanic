package leetcode

import (
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3,
			[]int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
	}

	for seq, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		for i := range test.nums1 {
			if test.nums1[i] != test.expect[i] {
				t.Errorf("case-%d: want: %v, but got: %v\n", seq+1,
					test.expect, test.nums1)
			}
		}
	}
}
