/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=30204
 *
 * [120] 三角形最小路径和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return -1
	}

	result := make([][]int, n)
	for i, _ := range result {
		result[i] = make([]int, n)
	}

	result[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		result[i][0] = result[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			result[i][j] = min(result[i-1][j], result[i-1][j-1]) + triangle[i][j]
		}
		result[i][i] = result[i-1][i-1] + triangle[i][i]
	}

	minResult := math.MaxInt32
	for _, r := range result[n-1] {
		minResult = min(minResult, r)
	}

	return minResult
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/

