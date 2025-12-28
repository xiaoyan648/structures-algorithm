/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30204
 *
 * [53] 最大子数组和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	preMax := nums[0]

	// dp[i] = max(dp[i-1] + nums[i], nums[i])
	for i := 1; i < len(nums); i++ {
		// max(preMax + nums[i], nums[i])
		if preMax > 0 {
			preMax += nums[i]
		} else {
			preMax = nums[i]
		}

		if preMax > max {
			max = preMax
		}
	}

	return max
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

