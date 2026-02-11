package practice

import "testing"

func TestScan(t *testing.T) {
	// 1. 当2个两部分，双指针法扫描两遍
	// 大小写字符串+数字 测试数据
	ints := []byte{'1', 'a', 'B', '3', 'c', 'D', '2', 'e', 'F', '4'}
	// ints2 := []byte{'a', 'B', 'c', 'D', 'e', 'F', '1', '3', '2', '4'}

	i, j := 0, 0
	for ; j < len(ints); j++ {
		if '0' <= ints[j] && ints[j] <= '9' {
			ints[i], ints[j] = ints[j], ints[i]
			i++
		}
	}
	printBytes(t, ints)

	j = i
	for ; j < len(ints); j++ {
		if 'A' <= ints[j] && ints[j] <= 'Z' {
			ints[i], ints[j] = ints[j], ints[i]
			i++
		}
	}

	// 转换为字符串数组打印
	printBytes(t, ints)
}

func printBytes(t *testing.T, bs []byte) {
	strs := make([]string, len(bs))
	for i, b := range bs {
		strs[i] = string(b)
	}
	t.Logf("%+v", strs)
}

func TestMappiing(t *testing.T) {
	m := [3][]byte{}

	ints := []byte{'1', 'a', 'B', '3', 'c', 'D', '2', 'e', 'F', '4'}

	for _, b := range ints {
		if '0' <= b && b <= '9' {
			m[0] = append(m[0], b)
		} else if 'A' <= b && b <= 'Z' {
			m[1] = append(m[1], b)
		} else if 'a' <= b && b <= 'z' {
			m[2] = append(m[2], b)
		}
	}

	result := make([]byte, 0, len(ints))
	for _, bs := range m {
		result = append(result, bs...)
	}

	printBytes(t, result)
}
