package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

https://leetcode.com/problems/two-sum
*/

// 只用一次循环来解决问题，用target减去当前数字，判断目标数字是否在之前遍历的时候出现过
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if _, exists := m[target-v]; exists {
			return []int{m[target-v], i}
		}
		m[v] = i
	}

	return nil
}

// 思路是先用map把nums中的数字保存起来，下一个循环遍历中判断想要的数是否在map中
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{i, m[target-v]}
		}
	}

	return nil
}
