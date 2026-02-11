package sort

import (
	"math/rand"
	"time"
)

// sort.Ints() // 标准库 快排、二分查找
// O(n^2)

func BubbleSort(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
}

func InsertSort(ints []int) {
	if len(ints) <= 1 {
		return
	}

	for i := 1; i < len(ints); i++ {
		temp := ints[i]

		j := i - 1
		for ; j >= 0 && temp < ints[j]; j-- {
			ints[j+1] = ints[j]
		}

		ints[j+1] = temp
	}
}

// MergeSort2 is an alternative implementation of MergeSort
// O(N*Log^n)

func MergeSort2(ints []int) {
	if len(ints) <= 1 {
		return
	}

	var coreSort func(i, j int)
	coreSort = func(i, j int) {
		if i >= j {
			return
		}

		mid := (i + j) / 2
		coreSort(i, mid)
		coreSort(mid+1, j)

		merge2(ints[i:j+1], ints[i:mid+1], ints[mid+1:j+1])
	}

	// actually run the recursive sort
	coreSort(0, len(ints)-1)
}

func merge2(ints, l, r []int) {
	temp := make([]int, 0, len(ints))

	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			temp = append(temp, l[i])
			i++
		} else {
			temp = append(temp, r[j])
			j++
		}
	}

	if i < len(l) {
		temp = append(temp, l[i:]...)
	}
	if j < len(r) {
		temp = append(temp, r[j:]...)
	}

	copy(ints, temp)
}

// MergeSort(i,j) = Merge(MergeSort(i,mid), MergeSort(mid+1,j))
// return when i >= j
func MergeSort(ints []int) {
	if len(ints) == 0 {
		return
	}

	mergeCore(ints, 0, len(ints)-1)
}

func mergeCore(ints []int, n, m int) {
	if n >= m {
		return
	}

	mid := (n + m) / 2
	mergeCore(ints, n, mid)
	mergeCore(ints, mid+1, m)
	// note, slice [n:m+1)
	merge(ints[n:m+1], ints[n:mid+1], ints[mid+1:m+1])
}

func merge(result []int, p1, p2 []int) {
	temp := make([]int, 0, len(result))

	// compare and set min value
	i, j := 0, 0
	for i < len(p1) && j < len(p2) {
		min := 0
		if p1[i] < p2[j] {
			min = p1[i]
			i++
		} else {
			min = p2[j]
			j++
		}
		temp = append(temp, min)
	}

	// set remain value in p1 or p2
	if i < len(p1) {
		temp = append(temp, p1[i:]...)
	} else if j < len(p2) {
		temp = append(temp, p2[j:]...)
	}

	// cover in result
	copy(result, temp)
}

// QuickSort O(n log n) average time complexity
func QuickSort2(ints []int) {
	if len(ints) <= 1 {
		return
	}

	var qsfunc func(i, j int)
	qsfunc = func(i, j int) {
		if i >= j {
			return
		}

		// 找到中间值，让分区跟均匀
		mid := (i + j) / 2
		if ints[i] > ints[j] {
			ints[i], ints[j] = ints[j], ints[i]
		}
		if ints[mid] < ints[i] {
			ints[mid], ints[i] = ints[i], ints[mid]
		}
		if ints[mid] > ints[j] {
			ints[mid], ints[j] = ints[j], ints[mid]
		}

		// 保存中间值到 j-1 位置，防止被交换覆盖
		ints[mid], ints[j] = ints[j], ints[mid]
		pivot := ints[j]

		// 双指针，当遇到小于 pivot 的值时，交换到左边
		left := i
		for right := i; right < j; right++ {
			if ints[right] < pivot {
				ints[right], ints[left] = ints[left], ints[right]
				left++
			}
		}
		// 最后将 pivot 放到中间位置
		ints[left], ints[j] = ints[j], ints[left]

		qsfunc(i, left-1)
		qsfunc(left+1, j)
	}

	qsfunc(0, len(ints)-1)
}

func QuickSort(ints []int) {
	if len(ints) == 0 {
		return
	}

	quickSortCore(ints, 0, len(ints)-1)
}

func quickSortCore(ints []int, n, m int) {
	if n >= m {
		return
	}

	pivot := partition(ints, n, m)
	quickSortCore(ints, n, pivot-1)
	quickSortCore(ints, pivot+1, m)
}

func partition(ints []int, n, m int) (pivot int) {
	pivotV := ints[m]

	i, j := n, n
	for i <= m && j <= m {
		if ints[j] < pivotV {
			j++
		} else {
			ints[j], ints[i] = ints[i], ints[j]
			i++
			j++
		}
	}

	if i == n {
		return n
	}
	return i - 1
}

// 快速排序O(n) 时间复杂度内求无序数组中的第 K 大元素。比如，4， 2， 5， 12， 3 这样一组数据，第 3 大元素就是 4。
func FindKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	rand.Seed(time.Now().UnixNano())
	return findKthLargestCore(nums, 0, len(nums)-1, k)
}

func findKthLargestCore(nums []int, n, m int, k int) int {
	// 找到中间点
	pivot := partitionDesc(nums, n, m)
	if k == pivot+1 {
		return nums[pivot]
	} else if k < pivot+1 {
		return findKthLargestCore(nums, n, pivot-1, k)
	} else {
		return findKthLargestCore(nums, pivot+1, m, k)
	}
}

// partition [n,m]获取第x大元素，返回其下标
func partitionDesc(nums []int, n, m int) int {
	randi := rand.Int()%(m-n+1) + n
	nums[randi], nums[m] = nums[m], nums[randi]
	pivotV := nums[m]

	i, j := n, n

	for ; j < m; j++ {
		if nums[j] < pivotV { // esc
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[m] = nums[m], nums[i]

	return i
}
