/*
 * @lc app=leetcode.cn id=1049 lang=golang
 * @lcpr version=30204
 *
 * [1049] 最后一块石头的重量 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lastStoneWeightII(stones []int) int {
	// 相减重量最小 => 找到两个子集，重量差最小 => 找到一个子集，重量接近 sum/2
	// 0-1背包：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	// 由于价值和重量相等且后面的数据依赖前面的，故简化为 dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
	sum := 0
	for _, n := range stones {
		sum += n
	}
	target := sum / 2

	dp := make([]int, target+1)
	for _, s := range stones {
		for j := target; j >= s; j-- {
			dp[j] = max(dp[j], dp[j-s]+s)
		}
	}

	return (sum - dp[target]) - dp[target]
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,4,1,8,1]\n
// @lcpr case=end

// @lcpr case=start
// [31,26,33,21,40]\n
// @lcpr case=end

*/

