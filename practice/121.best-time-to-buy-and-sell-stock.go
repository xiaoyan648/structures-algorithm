/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30204
 *
 * [121] 买卖股票的最佳时机
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxRes := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxRes = max(maxRes, prices[i]-minPrice)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

