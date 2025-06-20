package leetcode

import (
	"sort"
)

func partitionArray(nums []int, k int) int {
	ret := 1
	sort.Ints(nums)
	pnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-nums[pnt] > k {
			pnt = i
			ret++
		}
	}
	return ret
}
