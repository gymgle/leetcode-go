package leetcode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	maxLen := 0
	curLen := 0
	left := 0
	lookup := map[byte]bool{}
	for i := 0; i < n; i++ {
		curLen++
		for {
			if _, ok := lookup[s[i]]; ok {
				delete(lookup, s[left])
				curLen--
				left++
			} else {
				break
			}
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		lookup[s[i]] = true
	}
	return maxLen
}

