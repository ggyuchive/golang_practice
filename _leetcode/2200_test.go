package leetcode

import (
	"slices"
	"testing"
)

func Test2200(t *testing.T) {
	tests := []struct {
		nums   []int
		key    int
		k      int
		answer []int
	}{
		{[]int{3, 4, 9, 1, 3, 9, 5}, 9, 1, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 2, 2, 2, 2}, 2, 2, []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		actual := findKDistantIndices(tt.nums, tt.key, tt.k)
		if !slices.Equal(actual, tt.answer) {
			t.Errorf("findKDistantIndices(%v, %d, %d) = %v; want %v",
				tt.nums, tt.key, tt.k, actual, tt.answer)
		}
	}
}
