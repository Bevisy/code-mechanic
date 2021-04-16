package leetcode

//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？
//
//
//
//示例 1：
//
//输入：[1,0,2]
//输出：5
//解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。
//示例 2：
//
//输入：[1,2,2]
//输出：4
//解释：你可以分别给这三个孩子分发 1、2、1 颗糖果。
//第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/non-overlapping-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
func candy(ratings []int) int {
	size := len(ratings)
	candy := make([]int, size)
	// 遍历顺序：左-->右   左<右 ==> 右=左+1
	for i := 0; i < size-1; i++ {
		if ratings[i] < ratings[i+1] {
			candy[i+1] = candy[i] + 1
		}
	}

	//fmt.Printf("%v\n", candy)

	// 遍历顺序：右-->左   左>右 ==> 左=max(右+1, 左)
	for i := size - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candy[i-1] = max(candy[i]+1, candy[i-1])
		}
	}

	sum := size
	for _, v := range candy {
		sum += v
	}

	//fmt.Printf("%v\n", candy)

	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
