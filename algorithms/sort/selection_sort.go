package sort

//选择排序
//算法时间复杂度O(n^2)
func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				min = j
			}
		}
		tmp := array[min]
		array[min] = array[i]
		array[i] = tmp
	}

	return array
}

//func findMiniumNum(array []int) int {
//	ret := array[0]
//	index := 0
//	for k, v := range array {
//		if ret > v {
//			ret = v
//			index = k
//		}
//	}
//
//	return index
//}
