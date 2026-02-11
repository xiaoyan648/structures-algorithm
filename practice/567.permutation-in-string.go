/*
 * @lc app=leetcode.cn id=567 lang=golang
 * @lcpr version=30204
 *
 * [567] 字符串的排列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}

// @lc code=end

/*
// @lcpr case=start
// "eidbaooo"\n
// @lcpr case=end

// @lcpr case=start
// "eidboaoo"\n
// @lcpr case=end

*/

/*
func checkInclusion(s1 string, s2 string) bool {
	// 窗口大小=len(s1), 在 s2 这个窗口内的子串的元素数量是否和 s1 相等
	if len(s2) < len(s1) {
		return false
	}

	s1map := make(map[byte]int, min(len(s1), 26))
	s2map := make(map[byte]int, min(len(s2), 26))
	for i := 0; i < len(s1); i++ {
		s1map[s1[i]]++
		s2map[s2[i]]++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if isEqual(s1map, s2map) {
			return true
		}

		s2map[s2[i]]--
		s2map[s2[i+len(s1)]]++
	}

	return isEqual(s1map, s2map) // 最后一个窗口
}

func isEqual(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if v != m2[k] {
			return false
		}
	}

	return true
}
*/

