/*
 * @lc app=leetcode.cn id=438 lang=golang
 * @lcpr version=30204
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findAnagrams(s string, p string) []int {
	// Early return if p is longer than s
	if len(p) > len(s) {
		return []int{}
	}

	result := make([]int, 0)
	pmap := make(map[byte]int, len(p))
	smap := make(map[byte]int, len(p))

	for i := 0; i < len(p); i++ {
		pmap[p[i]]++
	}

	// init
	right := len(p)
	left := right - len(p)
	subString := s[left:right]
	for i := 0; i < len(subString); i++ {
		smap[subString[i]]++
	}
	if isEqual(pmap, smap) {
		result = append(result, left)
	}

	for right := len(p); right < len(s); right++ {
		// Remove leftmost character from window
		leftChar := s[right-len(p)]
		smap[leftChar]--
		if smap[leftChar] == 0 {
			delete(smap, leftChar)
		}
		// Add new rightmost character to window
		smap[s[right]]++

		if isEqual(pmap, smap) {
			result = append(result, right-len(p)+1)
		}
	}

	return result
}

func isEqual(pmap, smap map[byte]int) bool {
	if len(pmap) != len(smap) {
		return false
	}

	for k, v := range pmap {
		if smap[k] != v {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/

