/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30204
 *
 * [209] 长度最小的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	sum := 0

	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

/*
  1. 暴力破解法（原始代码）：
    - 外层循环：n 次
    - 内层循环：最坏情况下也是 n 次
    - 总复杂度：O(n²)
  2. 滑动窗口法：
    - 外层循环 right：0 到 n-1，只遍历一次 = O(n)
    - 内层循环 left：每次窗口满足条件时左指针右移，但 left 最多移动 n 次
    - 总复杂度：O(n)（因为 left 和 right 总共移动次数 ≤ 2n）

  为什么滑动窗口是 O(n)？
  关键在于两个指针都是单调递增的，都只会前进，不会后退，所以总移动次数是线性的。
*/

