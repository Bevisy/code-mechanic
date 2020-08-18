package leetcode

// 对撞指针
// 首尾分别 2 个指针，每次移动以后都分别判断长宽的乘积是否最大
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1 // max 记录面积最大值
	for start < end {                      // 当两个指针相遇时，循环终止
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}
