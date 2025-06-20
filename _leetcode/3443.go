package leetcode

import (
	"ggyuchive/math"
)

func maxDistance(s string, k int) int {
	p := "NSEW"
	ret := 0
	sum := make([]int, 4)
	for i := 0; i < len(s); i++ {
		for j := 0; j < 4; j++ {
			if p[j] == s[i] {
				sum[j]++
			}
		}
		cur := math.Abs(sum[0]-sum[1]) + math.Abs(sum[2]-sum[3])
		add := 2 * min(k, min(sum[0], sum[1])+min(sum[2], sum[3]))
		ret = max(ret, cur+add)
	}
	return ret
}
