/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30204
 *
 * [15] 三数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0, 3000)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i < len(nums) && i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			target := nums[i] + nums[left] + nums[right]
			if target < 0 {
				left++
			} else if target > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left-1] == nums[left] {
					left++
				}
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/

