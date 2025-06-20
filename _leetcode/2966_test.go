package leetcode

import (
	"ggyuchive/math"
	"testing"
)

// Test with special judge
func Test2966(t *testing.T) {
	tests := []struct {
		nums  []int
		k     int
		valid bool
	}{
		{[]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2, true},
		{[]int{2, 4, 2, 2, 5, 2}, 2, false},
		{[]int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, 14, true},
	}

	for _, tt := range tests {
		ret := divideArray(tt.nums, tt.k)

		if tt.valid && len(ret) == 0 {
			t.Errorf("Expected non-empty result for nums=%v, k=%d, got empty", tt.nums, tt.k)
		}
		if !tt.valid && len(ret) != 0 {
			t.Errorf("Expected empty result for nums=%v, k=%d, got %v", tt.nums, tt.k, ret)
		}

		// check all groups are of size 3 and satisfy the constraint
		for _, group := range ret {
			if len(group) != 3 {
				t.Errorf("Expected group of size 3, got %v", group)
			}
			for i := 0; i < 3; i++ {
				for j := i + 1; j < 3; j++ {
					if math.Abs(group[i]-group[j]) > tt.k {
						t.Errorf("Group %v violates k=%d constraint", group, tt.k)
					}
				}
			}
		}
	}
}
