package leetcode

func findKDistantIndices(nums []int, key int, k int) []int {
	ret := []int{}
	for i := 0; i < len(nums); i++ {
		for j := max(0, i-k); j <= min(len(nums)-1, i+k); j++ {
			if nums[j] == key {
				ret = append(ret, i)
				break
			}
		}
	}
	return ret
}
