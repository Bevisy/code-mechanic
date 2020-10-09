package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		expect int
	}{
		//{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{1, 2, 3}, 0, -1},
		//{[]int{}, 1, -1},
	}

	for _, test := range tests {
		got := binarySearch(test.array, test.target, 0, len(test.array)-1)
		if got != test.expect {
			t.Errorf("the expected answer is %d, but got %d\n", test.expect, got)
		}
	}
}

func TestIterBinarySearch(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		expect int
	}{
		{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{1, 2, 3}, 0, -1},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		got := iterBinarySearch(test.array, test.target, 0, len(test.array)-1)
		if got != test.expect {
			t.Errorf("the expected answer is %d, but got %d\n", test.expect, got)
		}
	}
}
