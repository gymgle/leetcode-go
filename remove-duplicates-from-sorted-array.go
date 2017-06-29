package leetcode

// Time complextiy : O(n)
// Space complexity : O(1)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
