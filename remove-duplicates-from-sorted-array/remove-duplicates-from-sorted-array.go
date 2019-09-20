package leetcode

/*
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

// Time complextiy : O(n)
// Space complexity : O(1)

/*
解题思路：
数组完成排序后，可以放置两个指针 i 和 j，其中 i 是慢指针，而 j 是快指针。
只要 nums[i] == nums[j]，增加 j 以跳过重复项。

当遇到 nums[i] != nums[j] 时，跳过重复项的运行已经结束，因此必须把它（nums[j]）的值复制到 nums[i+1]，
然后递增 i，接着将再次重复相同的过程，直到 j 到达数组的末尾为止。

在遍历 nums 数组时，可以用 for range 组合。
*/
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	i := 0
	for _, v := range nums[1:] {
		if v == nums[i] {
			continue
		}
		nums[i+1] = v
		i++
	}

	return i + 1
}
