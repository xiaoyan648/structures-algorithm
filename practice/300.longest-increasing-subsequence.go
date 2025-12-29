/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=30204
 *
 * [300] 最长递增子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLIS(nums []int) int {
	g := []int{}

	for _, x := range nums {
		i := sort.SearchInts(g, x)
		if i >= len(g) {
			g = append(g, x)
		} else {
			g[i] = x
		}
	}

	return len(g)
}

// @lc code=end

// 方法二：贪心 + 二分查找（优化到 O(n log n)）
// g[i] 表示长度为 i+1 的递增子序列中，末尾元素的最小值
// 关键思想：对于相同长度的递增子序列，末尾元素越小越好，因为更容易接上后续元素
func lengthOfLIS_Binary(nums []int) int {
	g := []int{} // g 数组保持严格递增

	for _, x := range nums {
		// 在 g 中二分查找第一个 >= x 的位置
		// sort.SearchInts 返回最小的 i 使得 g[i] >= x
		j := binarySearch(g, x)

		if j == len(g) {
			// 所有元素都 < x，说明 x 可以接在最长子序列后面
			g = append(g, x)
		} else {
			// 找到了 g[j] >= x，用 x 替换 g[j]
			// 这样长度为 j+1 的子序列末尾元素变得更小，更有利于后续扩展
			g[j] = x
		}
	}

	return len(g)
}

// 二分查找：返回第一个 >= target 的位置
func binarySearch(g []int, target int) int {
	left, right := 0, len(g)
	for left < right {
		mid := (right + left) / 2
		if g[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/*
详细解释：

1. 为什么 g 数组是严格递增的？
   - 假设 g[i] >= g[i+1]
   - 那么长度为 i+2 的子序列去掉最后一个元素，得到长度为 i+1 的子序列
   - 这个子序列的末尾必然 < g[i+1]，与 g[i] 的定义矛盾
   - 所以 g 必然严格递增

2. 为什么用 x 替换 g[j]？
   - g[j] >= x，说明长度为 j+1 的子序列末尾可以是 x
   - 用更小的 x 替换 g[j]，使得后续元素更容易满足递增条件
   - 这是贪心策略：保持每个长度的子序列末尾尽可能小

3. 时间复杂度分析：
   - 外层循环 O(n)
   - 内层二分查找 O(log n)
   - 总时间复杂度 O(n log n)

示例演示：nums = [10, 9, 2, 5, 3, 7, 101, 18]

步骤        x    j    操作          g
初始        -    -    -            []
处理10      10   0    append       [10]
处理9       9    0    g[0]=9       [9]         // 用9替换10，长度1的子序列末尾更小
处理2       2    0    g[0]=2       [2]         // 用2替换9，更小
处理5       5    1    append       [2,5]       // 5>2，可以扩展
处理3       3    1    g[1]=3       [2,3]       // 用3替换5，长度2的子序列末尾更小
处理7       7    2    append       [2,3,7]     // 7>3，可以扩展
处理101     101  3    append       [2,3,7,101] // 扩展到长度4
处理18      18   3    g[3]=18      [2,3,7,18]  // 用18替换101，末尾更小

最终答案：len(g) = 4，对应子序列 [2,3,7,18] 或 [2,3,7,101] 等
*/

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

