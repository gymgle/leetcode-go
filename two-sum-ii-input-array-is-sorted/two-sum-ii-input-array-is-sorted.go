package leetcode

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2

https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
*/

// 使用两个指针分别指向数组的第一个位置和最后一个位置向中间夹逼

func twoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1
	for lp < rp {
		if numbers[lp]+numbers[rp] < target {
			lp++
		} else if numbers[lp]+numbers[rp] > target {
			rp--
		} else {
			return []int{lp + 1, rp + 1}
		}
	}
	return nil
}
