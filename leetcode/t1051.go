package leetcode

import "sort"

func heightChecker(heights []int) int {
	var count int
	// 初始化固定长度数组
	s := make([]int, len(heights))
	// 深拷贝
	copy(s, heights)
	// 排序
	sort.Ints(s)

	for i := range s {
		if s[i] != heights[i] {
			count++
		}
	}

	return count
}
