/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=30204
 *
 * [198] 打家劫舍
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {
	// dp[i] 表示到达 i 时最大偷窃值
	// dp[i] = max(dp[i-1], dp[i-2]+dp[i])
	if len(nums) == 1 {
		return nums[0]
	}

	p2, p1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(p1, p2+nums[i])
		p2 = p1
		p1 = tmp
	}

	return p1
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
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/

