package dp

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestMinEditCountRecursive(t *testing.T) {

	fmt.Println((3 - 1000) % 2)
	a := "mitcmu"
	b := "mtacnu"
	alen, blen := 6, 6
	minEditCount := math.MaxInt

	var minEditFunc func(i, j, mec int)
	minEditFunc = func(i, j, mec int) {
		if i == alen || j == blen {
			var min int
			if i < alen {
				min = mec + (alen - i)
			} else if j < blen {
				min = mec + (blen - j)
			} else {
				min = mec
			}

			if min < minEditCount {
				minEditCount = min
			}
			return
		}

		if a[i] == b[j] {
			minEditFunc(i+1, j+1, mec)
		} else {
			minEditFunc(i+1, j, mec+1)
			minEditFunc(i, j+1, mec+1)
			minEditFunc(i+1, j+1, mec+1)
		}
	}

	minEditFunc(0, 0, 0)

	fmt.Println(minEditCount)
}

func TestMinEditCountDP(t *testing.T) {
	a := "mitcmu"
	b := "mtacnu"
	alen, blen := 6, 6

	firstValue := 0
	if a[0] != b[0] {
		firstValue = 1
	}

	minEditGraph := [6][6]int{}
	// 设置第一层参数
	for i := firstValue; i < alen; i++ {
		minEditGraph[0][i] = i
	}
	for j := firstValue; j < blen; j++ {
		minEditGraph[j][0] = j
	}

	min := func(i, j, n int) int {
		min := i
		if j < min {
			min = j
		}
		if n < min {
			min = n
		}
		return min
	}

	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			minEditGraph[i][j] = min(minEditGraph[i-1][j], minEditGraph[i][j-1], minEditGraph[i-1][j-1])
			if a[i] != b[j] {
				minEditGraph[i][j] += 1
			}
		}
	}

	fmt.Println(minEditGraph[5][5])
}

func TestMaxSubString(t *testing.T) {
	// 最长公共子串
	// f(i,j) = max(f(i-1,j-1) ?+1, f(i-1,j), f(i,j-1))

	a := "mitcmu"
	b := "mtacnu"
	alen, blen := 6, 6
	maxSubString := [6][6]int{}
	if a[0] == b[0] {
		maxSubString[0][0] = 1
	}
	for i := 1; i < alen; i++ {
		maxSubString[i][0] = i
	}
	for j := 1; j < blen; j++ {
		maxSubString[0][j] = j
	}

	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			if a[i] == b[j] {
				maxSubString[i][j] = max(maxSubString[i-1][j-1]+1, maxSubString[i-1][j], maxSubString[i][j-1])
			} else {
				maxSubString[i][j] = max(maxSubString[i-1][j-1], maxSubString[i][j-1], maxSubString[i-1][j])
			}
		}
	}
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	return b
}

/*
一个数字序列包含 n 个不同的数字，如何求出这个序列中的最长递增子序列长度？比如 2, 9, 3, 6, 5, 1, 7 这样一组数字序列，它的最长递增子序列就是 2, 3, 5, 7，所以最长递增子序列的长度是 4。
*/

func TestMaxIncreseSeq(t *testing.T) {
	data := []int{2, 9, 3, 6, 5, 1, 7}
	dataLen := len(data)
	maxSeqNum := make([]int, dataLen)

	maxValue := 0
	maxSeqNum[0] = 1
	for i := 1; i < dataLen; i++ {
		itemMaxValue := 0
		for j := i - 1; j >= 0; j-- {
			if data[i] > data[j] && itemMaxValue < maxSeqNum[j]+1 {
				itemMaxValue = maxSeqNum[j] + 1
			}
		}
		maxSeqNum[i] = itemMaxValue

		if itemMaxValue > maxValue {
			maxValue = itemMaxValue
		}
	}

	println(maxValue)
}

func TestMinPathSum(t *testing.T) {
	r := minPathSum([][]int{{1, 2, 3}, {4, 5, 6}})
	print(r)
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	minDist := make([][]int, m)
	for i := range minDist {
		minDist[i] = make([]int, n)
	}

	minDist[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		minDist[i][0] = minDist[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		minDist[0][j] = minDist[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			minDist[i][j] = min(minDist[i-1][j], minDist[i][j-1]) + grid[i][j]
		}
	}

	return minDist[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestMinCoinNumForTarget(t *testing.T) {
	target := 9
	coins := []int{1, 3, 5}
	sort.Ints(coins)
	cache := make([]int, target)

	println(minCoinNums(coins, target, cache))
}

func minCoinNums(coins []int, target int, cache []int) int {
	if cache[target-1] != 0 {
		return cache[target-1]
	}

	validCoins := coins
	for i, c := range coins {
		if target == c {
			return 1
		}
		if target < c {
			validCoins = coins[:i]
			break
		}
	}
	if len(validCoins) == 0 {
		return 0
	}

	minNum := math.MaxInt32
	for _, c := range validCoins {
		if target-c > 0 {
			if n := minCoinNums(validCoins, target-c, cache); n < minNum {
				minNum = n
			}
		}
	}

	result := minNum + 1
	cache[target-1] = result
	return result
}
