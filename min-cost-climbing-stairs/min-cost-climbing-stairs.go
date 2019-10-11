/**
 * 746. Min Cost Climbing Stairs
 * https://leetcode.com/problems/min-cost-climbing-stairs/
 */

package leetcode

// 使用动态规划的思路：假设 f(i) 表示从第 i 层到顶层的最小花费，那么 f(i) = cost[i] + min(f(i+1), f(i+2))
func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f := cost[i] + min(f1, f2)
		f1, f2 = f2, f
	}

	return min(f1, f2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
