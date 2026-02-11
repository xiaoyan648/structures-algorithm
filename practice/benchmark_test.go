package practice

import (
	"fmt"
	"testing"
)

// 暴力破解法（原始）
func minSubArrayLenBrute(target int, nums []int) int {
	minLen := math.MaxInt

	for l := 0; l <= len(nums)-1; l++ {
		sum, length := 0, 0
		for r := l; r <= len(nums)-1; r++ {
			sum += nums[r]
			length++
			if sum >= target && minLen > length {
				minLen = length
				break
			}
		}
	}

	if minLen == math.MaxInt {
		minLen = 0
	}

	return minLen
}

// 滑动窗口法（优化）
func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
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

// 基准测试函数
func benchmarkAlgorithm(algorithm func(int, []int) int, target int, nums []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithm(target, nums)
	}
}

func TestPerformanceComparison(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
	}{
		{
			name:   "Case 1: 短数组，快速找到解",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3}, // [2,3,1,2,4,3]，答案 2 ([4,3])
		},
		{
			name:   "Case 2: 长数组，需要全遍历",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "Case 3: 大数组，重复元素",
			target: 100,
			nums:   make([]int, 1000), // 1000个0
		},
		{
			name:   "Case 4: 递减数组",
			target: 8,
			nums:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:   "Case 5: 无解情况",
			target: 100,
			nums:   []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 运行基准测试
			fmt.Printf("\n%s:\n", tt.name)
			fmt.Printf("输入: target=%d, nums长度=%d\n", tt.target, len(tt.nums))

			// 测试结果正确性
			result1 := minSubArrayLenBrute(tt.target, tt.nums)
			result2 := minSubArrayLen(tt.target, tt.nums)

			if result1 != result2 {
				t.Errorf("结果不匹配: 暴力破解=%d, 滑动窗口=%d", result1, result2)
			}

			fmt.Printf("结果: %d\n", result1)
		})
	}
}
