package main

import (
	"fmt"
	"math"
	"strings"
	"time"
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

// 性能测试函数
func benchmark(name string, algorithm func(int, []int) int, target int, nums []int) {
	start := time.Now()
	result := algorithm(target, nums)
	duration := time.Since(start)

	fmt.Printf("\n%s:\n", name)
	fmt.Printf("  输入: target=%d, nums长度=%d\n", target, len(nums))
	fmt.Printf("  结果: %d\n", result)
	fmt.Printf("  耗时: %v\n", duration)
	fmt.Printf("  指针移动次数分析:\n")

	// 简单的复杂度分析
	if name == "暴力破解法" {
		fmt.Printf("    - 外层循环: %d 次\n", len(nums))
		fmt.Printf("    - 内层循环: 约 %d 次 (总和)\n", len(nums)*(len(nums)+1)/2)
		fmt.Printf("    - 时间复杂度: O(n²)\n")
	} else {
		fmt.Printf("    - right指针: %d 次 (0→%d)\n", len(nums), len(nums)-1)
		fmt.Printf("    - left指针: %d 次 (0→%d)\n", len(nums), len(nums)-1)
		fmt.Printf("    - 总移动次数: ≤ %d 次\n", 2*len(nums))
		fmt.Printf("    - 时间复杂度: O(n)\n")
	}
}

func main() {
	testCases := []struct {
		name   string
		target int
		nums   []int
	}{
		{
			name:   "Case 1: 短数组，快速找到解",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3}, // 答案 2 ([4,3])
		},
		{
			name:   "Case 2: 长数组，需要全遍历",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "Case 3: 大数组（5000元素）",
			target: 5000,
			nums:   make([]int, 5000), // 5000个0，然后加一些大数
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

	fmt.Println("=== 两种算法性能对比 ===")
	fmt.Println("目标：展示滑动窗口相比暴力破解的优化效果\n")

	for _, tt := range testCases {
		benchmark("暴力破解法", minSubArrayLenBrute, tt.target, tt.nums)
		fmt.Println()
		benchmark("滑动窗口法", minSubArrayLen, tt.target, tt.nums)
		fmt.Println(strings.Repeat("-", 60))
	}

	// 演示复杂度的关键差异
	fmt.Println("\n=== 复杂度分析总结 ===")
	sizes := []int{10, 100, 1000}
	fmt.Println("\nn的大小 vs 操作次数对比:")
	fmt.Println("n\t暴力破解\t滑动窗口\t优化倍数")
	fmt.Println(strings.Repeat("-", 50))

	for _, n := range sizes {
		bruteOps := n * n          // O(n²)
		windowOps := 2 * n         // O(n)
		ratio := bruteOps / windowOps
		fmt.Printf("%d\t%d\t\t%d\t\t%d倍\n", n, bruteOps, windowOps, ratio)
	}
}
