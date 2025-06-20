package leetcode

import (
	"testing"
)

func Test2294(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		answer int
	}{
		{[]int{3, 6, 1, 2, 5}, 2, 2},
		{[]int{1, 2, 3}, 1, 2},
		{[]int{2, 2, 4, 5}, 0, 3},
	}

	for _, tt := range tests {
		actual := partitionArray(tt.nums, tt.k)
		if actual != tt.answer {
			t.Errorf("partitionArray(%v, %d) = %d; want %d",
				tt.nums, tt.k, actual, tt.answer)
		}
	}
}
