package leetcode

/**
 * 242. Valid Anagram
 *
 * Given two strings s and t , write a function to determine if t is an anagram of s.
 *
 * Example 1:
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * Example 2:
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 * https://leetcode.com/problems/valid-anagram/
 */

// 该方法最容易想到，使用两个map分别保存两个字符串中的字符和出现次数，最后比较两个map是否相同
// 三次遍历，两个map空间
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, v := range s {
		m1[v]++
	}

	for _, v := range t {
		m2[v]++
	}

	return isEqual(m1, m2)
}

func isEqual(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

// 方法二，使用一个map，第一次遍历增加计数，第二次遍历减少计数，最后比较map中是否含有非0的value
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}
	for _, c := range t {
		counts[c]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}
