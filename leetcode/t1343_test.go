package leetcode

import "testing"

func TestNumOfSubarrays(t *testing.T) {
	arrays := []struct {
		input     []int
		k         int
		threshold int
		expected  int
	}{
		{[]int{1, 1, 1, 1, 1}, 1, 0, 5},
		{[]int{1}, 1, 1, 0},
		{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4, 3},
	}

	for _, test := range arrays {
		got := numOfSubarrays(test.input, test.k, test.threshold)
		if test.expected != got {
			t.Errorf("Inout: %v, Expected: %d, Got: %d\n", test.input,
				test.expected, got)
		}
	}
}
