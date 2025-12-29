/*
 * @lc app=leetcode.cn id=1143 lang=golang
 * @lcpr version=30204
 *
 * [1143] 最长公共子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := [2][]int{
		make([]int, len(text2)+1),
		make([]int, len(text2)+1),
	}

	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				dp[(i+1)%2][j+1] = dp[i%2][j] + 1
			} else {
				dp[(i+1)%2][j+1] = max(dp[(i+1)%2][j], dp[i%2][j+1])
			}
		}
	}

	return dp[len(text1)%2][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "abcde"\n"ace"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"def"\n
// @lcpr case=end

*/

/*
1. 递归 + 缓存剪枝
func longestCommonSubsequence(text1 string, text2 string) int {
	var dfs func(i, j int) int
	mem := make([][]int, len(text1))
	for i := range mem {
		mem[i] = make([]int, len(text2))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if mem[i][j] >= 0 {
			return mem[i][j]
		}

		var v int
		if text1[i] == text2[j] {
			v = dfs(i-1, j-1) + 1
		} else {
			v = max(dfs(i-1, j), dfs(i, j-1))
		}

		mem[i][j] = v
		return v
	}

	return dfs(len(text1)-1, len(text2)-1)
}
*/