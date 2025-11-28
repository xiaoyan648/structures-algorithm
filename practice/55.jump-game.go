/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=30204
 *
 * [55] 跳跃游戏
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// 正着跳
func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}

		maxReach = max(i+nums[i], maxReach)

		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 反着跳
// func canJump(nums []int) bool {
// 	target := len(nums) - 1 // target 表示当前需要调到的最后一位

// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if i+nums[i] >= target {
// 			target = i
// 		}
// 	}

// 	return target == 0
// }

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/

