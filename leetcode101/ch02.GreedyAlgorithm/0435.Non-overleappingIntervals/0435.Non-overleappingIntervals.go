package leetcode

import "sort"

//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//
//注意:
//
//可以认为区间的终点总是大于它的起点。
//区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
//示例 1:
//
//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//
//输出: 1
//
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
//示例 2:
//
//输入: [ [1,2], [1,2], [1,2] ]
//
//输出: 2
//
//解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
//示例 3:
//
//输入: [ [1,2], [2,3] ]
//
//输出: 0
//
//解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-overlapping-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(dataType(intervals))
	removed, prev := 0, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 判断区间是否重叠
		if intervals[i][0] < prev {
			removed++
		} else {
			prev = intervals[i][1]
		}
	}
	return removed
}

// 二维整型数组，第二数组元素末尾大小排序
type dataType [][]int

func (p dataType) Len() int {
	return len(p)
}

func (p dataType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p dataType) Less(i, j int) bool {
	return p[i][len(p[i])-1] < p[j][len(p[j])-1]
}
