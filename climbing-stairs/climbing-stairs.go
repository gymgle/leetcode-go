/**
 * 70. Climbing Stairs
 * https://leetcode.com/problems/climbing-stairs/
 */

package leetcode

// 使用动态规划的思路：从最高台阶反向求解，状态转移方程为 dp(n) = dp(n-1) + dp(n-2)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	a := 1
	b := 2

	var c int
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}
