package sort

//插入排序
//算法时间复杂度O(n^2)
func insertionSort(array []int) []int {
	for currentIndex := 1; currentIndex < len(array); currentIndex++ {
		temp := array[currentIndex]
		iterator := currentIndex
		//已排序的数组，从后往前检索，直到找到需插入数据的合适位置
		for ; iterator > 0 && array[iterator-1] >= temp; iterator-- {
			array[iterator] = array[iterator-1]
		}
		array[iterator] = temp
	}
	return array
}
