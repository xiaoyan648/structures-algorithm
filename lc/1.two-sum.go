/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30204
 *
 * [1] 两数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, x := range nums {
		if j, ok := hashMap[target-x]; ok {
			return []int{j, i}
		}
		hashMap[x] = i
	}
	return nil
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/

