/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=30204
 *
 * [64] 最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	minDist := make([][]int, m)
	for i, _ := range minDist {
		minDist[i] = make([]int, n)
	}

	minDist[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		minDist[i][0] = minDist[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		minDist[0][j] = minDist[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			minDist[i][j] = min(minDist[i-1][j], minDist[i][j-1]) + grid[i][j]
		}
	}

	return minDist[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/

