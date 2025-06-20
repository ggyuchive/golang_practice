package leetcode

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	n := len(nums)
	m := n / 3
	ret := make([][]int, m)
	ret2 := [][]int{}
	sort.Ints(nums)

	pnt := 0
	for i := 0; i < m; i++ {
		group := make([]int, 3)
		for j := 0; j < 3; j++ {
			group[j] = nums[pnt]
			pnt++
		}
		if group[2]-group[0] > k {
			return ret2
		}
		ret[i] = group
	}
	return ret
}
