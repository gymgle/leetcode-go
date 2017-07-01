package leetcode

/*
Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.

https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	index := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[index] == nums[i] && count == 0 {
			continue
		} else {
			if nums[index] == nums[i] {
				count--
			} else {
				count = 1
			}

			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
