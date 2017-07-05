package leetcode

/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

For example,
Given [100, 4, 200, 1, 3, 2],
The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.

Your algorithm should run in O(n) complexity.

https://leetcode.com/problems/longest-consecutive-sequence
*/

func longestConsecutive(nums []int) int {
	var m = make(map[int]bool)
	max := 0
	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		pre := v - 1
		next := v + 1
		for _, visited := m[pre]; visited; _, visited = m[pre] {
			m[pre] = false
			pre--
		}

		for _, visited := m[next]; visited; _, visited = m[next] {
			m[next] = false
			next++
		}

		if max < next-pre-1 {
			max = next - pre - 1
		}
	}
	return max
}
