/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=30204
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	// 只要在最低点买、最高点卖即可
	// 检查是否递增，如果是，则在第一次增加时买入；
	// 检查是否递减，如果是，则在第一次减少时卖出；

	if len(prices) <= 1 {
		return 0
	}

	isBuy := false // 未买入、已买入
	buyMoneny := 0
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] && !isBuy {
			isBuy = true
			buyMoneny = prices[i-1]
		}
		if prices[i-1] > prices[i] && isBuy {
			isBuy = false
			result += prices[i-1] - buyMoneny
		}
	}

	if isBuy {
		result += prices[len(prices)-1] - buyMoneny
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

