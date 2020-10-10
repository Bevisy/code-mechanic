package sort

// BubbleSort 冒泡排序. data必须实现sort包中的Interface接口
//func BubbleSort(data sort.Interface) {
//	n := data.Len()
//	for i := 0; i < n-1; i++ {
//		isChanged := false
//		for j := 0; j < n-1-i; j++ {
//			if data.Less(j, j+1) {
//				data.Swap(j, j+1)
//				isChanged = true
//			}
//		}
//		if !isChanged {
//			break
//		}
//	}
//}

// declare a array
// this array must implenet sort.Inerface
//func main() {
//	data := sort.IntSlice{22, 34, 3, 40, 18, 4}
//	BubbleSort(data)
//}

// 方法二：针对 int 数组
func bubbleSort(array []int) []int {
	//循环比较次数
	for i := 0; i < len(array)-1; i++ {
		isChanged := false
		//比较相邻两数，符合条件则交换
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
				isChanged = true
			}
		}

		if !isChanged {
			break
		}
	}

	return array
}
