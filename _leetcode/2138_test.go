package leetcode

import (
	"slices"
	"testing"
)

func Test2138(t *testing.T) {
	tests := []struct {
		s      string
		k      int
		fill   byte
		answer []string
	}{
		{"abcdefghi", 3, 'x', []string{"abc", "def", "ghi"}},
		{"abcdefghij", 3, 'x', []string{"abc", "def", "ghi", "jxx"}},
	}

	for _, tt := range tests {
		actual := divideString(tt.s, tt.k, tt.fill)
		if !slices.Equal(actual, tt.answer) {
			t.Errorf("divideString(%s, %d, %d) = %v; want %v",
				tt.s, tt.k, tt.fill, actual, tt.answer)
		}
	}
}
