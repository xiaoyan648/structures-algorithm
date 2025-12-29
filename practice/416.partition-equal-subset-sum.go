/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=30204
 *
 * [416] 分割等和子集
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canPartition(nums []int) bool {
	// 0-1背包：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	// 由于价值和重量相等且后面的数据依赖前面的，故简化为 dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
	// 我们只需要单个一维数组 dp[j] 即可，但是要注意从后往前遍历，放在前面变量变化应用后面结果
	// 因为是一维数组所以我们不需要考虑 j < nums[i] 的情况, 前置的迭代已经赋值。
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/

/*

// 0-1背包：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	// 目前要找到两个相同大小的序列代表 target（背包价值） = sum/2
	// 因为 重量==价值 的，dp[w]（表示w承重的背包） = v（最大价值）, v <= w, 装满：dp[target] == target

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= len(nums); i++ {
		nindex := i - 1
		for j := 1; j <= target; j++ {
			if j >= nums[nindex] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[nindex]]+nums[nindex])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
		if dp[i][target] == target {
			return true
		}
	}

	return false

// 0-1背包：内存优化：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	// 由于价值和重量相等，故简化为 dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
	// if dp[j] == target 则表示找到了
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 { // 奇数无法分成两份
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}

	return dp[target] == target

// 0-1背包：dp[i][w] = max(dp[i-1][w], dp[i-1][w-weight[i]] + value[i])
	// 由于我们只需要知道有能凑出 target 的秩序列，所以递推公式的含义可以是：
	// dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]], 表示前 i 个数能否凑出和为 j子序列
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 { // 奇数无法分成两份
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	minTarget := 0

	for _, num := range nums {
		minTarget = min(minTarget+num, target)
		for j := minTarget; j >= num; j-- {
			dp[j] = dp[j-num] || dp[j]
		}
		if dp[target] {
			return true
		}
	}

	return dp[target]
// 效率优化：bit位计算
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 { // 奇数无法分成两份
		return false
	}

	target := sum / 2
	n, m := big.NewInt(1), big.NewInt(0)
	for _, num := range nums {
		n.Or(n, m.Lsh(n, uint(num)))
	}
	return n.Bit(target) == 1
*/