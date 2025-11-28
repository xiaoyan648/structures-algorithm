/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=30204
 *
 * [70] 爬楼梯
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	c1, c2, r := 1, 2, 0

	for i := 3; i <= n; i++ {
		r = c1 + c2
		c1 = c2
		c2 = r
	}

	return r
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

