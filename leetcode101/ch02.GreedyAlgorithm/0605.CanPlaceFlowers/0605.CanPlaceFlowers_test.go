package leetcode

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 0}, 1, false},
	}

	for _, test := range tests {
		got := canPlaceFlowers(test.flowerbed, test.n)
		if got != test.expected {
			t.Errorf("expect %t, but got %t\n", test.expected, got)
		}
	}
}
