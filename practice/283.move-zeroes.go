/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30204
 *
 * [283] 移动零
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

