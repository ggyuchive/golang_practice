package leetcode

func divideString(s string, k int, fill byte) []string {
	n, grp := len(s), len(s)/k
	if n%k > 0 {
		grp++
	}
	ret := make([]string, grp)
	i := 0
	for _, c := range s {
		ret[i] += string(c)
		if len(ret[i]) == k {
			i++
		}
	}
	for len(ret[grp-1]) < k {
		ret[grp-1] += string(fill)
	}
	return ret
}
