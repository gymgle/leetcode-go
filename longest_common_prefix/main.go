package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
	}
	return prefix
}

func lcp(s1, s2 string) string {
	length := min(len(s1), len(s2))
	i := 0
	for i < length && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
