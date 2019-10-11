/**
 * 63. Unique Paths II
 * https://leetcode.com/problems/unique-paths-ii/
 */

package leetcode

// 动态规划的思路：考虑机器人从左上角想右下角移动，矩阵的第一行和第一列的每个位置，都只有一种走法
// 对于有障碍的位置，该位置的走法有0种。
// 状态定义：p[i][j] 表示 i 行 j 列的走法数量
// 状态转移方程：
//     如果有障碍，p[i][j] = 0
//     特殊考虑第一行和第一列：p[0][0] = 1 - obstacleGridp[0][0];
//                         p[i][0] = p[i-1][0]
//                         p[0][j] = p[0][j-1]
//     其他位置：
//     obstacleGrid[i][j] == 0 时，p[i][j] = p[i-1][j] + p[i][j-1]
//     obstacleGrid[i][j] == 0 时，p[i][j] = 0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	p := make([][]int, m)
	for i := 0; i < m; i++ {
		p[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				p[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					p[i][j] = 1 - obstacleGrid[i][j]
				} else if i == 0 {
					p[i][j] = p[i][j-1]
				} else if j == 0 {
					p[i][j] = p[i-1][j]
				} else {
					p[i][j] = p[i-1][j] + p[i][j-1]
				}
			}
		}
	}

	return p[m-1][n-1]
}
