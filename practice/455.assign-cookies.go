/*
 * @lc app=leetcode.cn id=455 lang=golang
 * @lcpr version=30204
 *
 * [455] 分发饼干
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}

	return i
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[1,2,3]\n
// @lcpr case=end

*/

