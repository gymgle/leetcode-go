package leetcode

/*
80. Remove Duplicates from Sorted Array II

Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.

https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/

/*
解题思路：在上一道题的基础上加一个计数器 count
count 初始化为 1 表示最多重复 1 次。当 count 初始化为 0 时，退化为上一题。
*/
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	i := 0
	count := 1
	for _, v := range nums[1:] {
		if nums[i] == v && count == 0 {
			continue
		} else {
			if nums[i] == v {
				count--
			} else {
				count = 1
			}

			i++
			nums[i] = v
		}
	}
	return i + 1
}
