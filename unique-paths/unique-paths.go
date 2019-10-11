/**
 * 62. Unique Paths
 * https://leetcode.com/problems/unique-paths/
 */

package leetcode

// 动态规划的思路：考虑机器人从左上角想右下角移动，矩阵的第一行和第一列的每个位置，都只有一种走法
// 剩下的位置 p[i][j] 表示 i 行 j 列的走法数量(定义状态)，则状态转移方程是：
// p[i][j] = p[i-1][j] + p[i][j-1]
func uniquePaths(m int, n int) int {
	// 初始化矩阵, go 无法定义动态大小的数组, 不可以写成 p := [m][n]int{}
	p := make([][]int, m) // 由 m 个 slice 组成
	for i := range p {
		p[i] = make([]int, n) // 每行是一个包含 n 个数字的 slice
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 对于第一行和第一列的位置，只有一种走法
			if i == 0 || j == 0 {
				p[i][j] = 1
			} else {
				p[i][j] = p[i-1][j] + p[i][j-1]
			}
		}
	}

	return p[m-1][n-1]
}
