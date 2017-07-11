/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

https://leetcode.com/problems/3sum

*/

// 可以参照维基百科上的二次复杂度算法

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := 0 - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			hit := target - nums[j]
			if hit < nums[j] {
				break
			}
			for k := j + 1; k < len(nums) && hit >= nums[k]; k++ {
				if hit == nums[k] {
					tmp := []int{nums[i], nums[j], nums[k]}
					ret = append(ret, tmp)
					break
				}
			}
		}
	}
	return ret
}
