package leetcode

import (
	"testing"
)

func Test3443(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		answer int
	}{
		{"NWSE", 1, 3},
		{"NSWWEW", 3, 6},
	}

	for _, tt := range tests {
		actual := maxDistance(tt.s, tt.k)
		if actual != tt.answer {
			t.Errorf("maxDistance(%s, %d) = %d; want %d",
				tt.s, tt.k, actual, tt.answer)
		}
	}
}
