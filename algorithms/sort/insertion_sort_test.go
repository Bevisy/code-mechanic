package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		got := insertionSort(test.input)
		for i := range test.expect {
			if test.expect[i] != got[i] {
				t.Errorf("expect %v, but got %v\n", test.expect, got)
				break
			}
		}
	}
}
