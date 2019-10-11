/**
 * 53. Maximum Subarray
 * https://leetcode.com/problems/maximum-subarray/
 */

package leetcode

// 动态规划的思路
// 定义状态: f(i) 表示以第 i 个数结尾的子数组最大和
// 求: max(f(i)), 0 <= i < n
// 状态转移方程:
//   如果 f(i-1)<0, f(i) = nums[i]
//   如果 f(i-1)>0, f(i) = nums[i] + f(i-1)

// 代码可以写成这样：
/*
func maxSubArray(nums []int) int {
	lens := len(nums)
	if lens == 0 {
		return 0
	}

	max := nums[0]
	f1 := nums[0]
	for i := 1; i < lens; i++ {
		f2 := nums[i]
		if f1 > 0 {
			f2 = nums[i] + f1
		}

        f1 = f2

		// 记录出现过的最大值
		if f2 > max {
			max = f2
		}
	}

	return max
}
*/

// 还可以继续精简，去掉 f1、f2 变量，统一为 curMax，使用 for range 遍历数组
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	curSum := 0
	for _, v := range nums {
		if curSum > 0 {
			curSum += v
		} else {
			curSum = v
		}

		// 记录出现过的最大值
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}
