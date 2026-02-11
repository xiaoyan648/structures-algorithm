/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 3
 * [3] 无重复字符的最长子串
*/

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	cache := make(map[byte]int, 0)
	maxLen := math.MinInt

	for ; right < len(s); right++ {
		if lastIndex, ok := cache[s[right]]; ok {
			left = max(lastIndex+1, left)
		}
		cache[s[right]] = right

		maxLen = max(right-left+1, maxLen)
	}

	if maxLen == math.MinInt {
		return 0
	}

	return maxLen
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/

