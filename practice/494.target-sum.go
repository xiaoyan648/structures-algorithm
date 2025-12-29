/*
 * @lc app=leetcode.cn id=494 lang=golang
 * @lcpr version=30204
 *
 * [494] 目标和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0 
	for _, n := range nums {
		sum+=n
	}
	if target > sum ||(sum+target) < 0 ||(sum+target) % 2 != 0 { // P > 0 P%
		return 0
	}

	p := (sum+target)/2

	dp := make([]int, p+1)
	dp[0] = 1

	for _, num := range nums {
		for j := p; j >= num; j-- {
			dp[j] += dp[j-num] // dp[j] num[:i] 能达成的 + dp[j-num] = num[:i+1] 能达成的数量
		}
	}

	return dp[p]
}

// @lc code=end

/*
DP解法推导：
1. 数学转换：
   设P为正数和，N为负数和，total为总和
   P - N = total
   P + N = target
   => 2P = target + total
   => P = (target + total) / 2
   问题转化为：寻找子集和为P的方案数

2. DP状态转移：
   dp[i][j] = 前i个元素组成和为j的方案数
   dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
*/

// DP解法
func findTargetSumWaysDP(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	// 检查可行性：(target + total) 必须为偶数且非负
	if target > total || (target+total)%2 != 0 {
		return 0
	}

	P := (target + total) / 2

	// dp[j] = 和为j的方案数
	dp := make([]int, P+1)
	dp[0] = 1

	for _, num := range nums {
		// 逆序遍历避免重复使用同一个元素
		for j := P; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[P]
}

// @lcpr case=start
// [1,1,1,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/

