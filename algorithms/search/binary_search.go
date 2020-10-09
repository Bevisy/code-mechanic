package search

//复杂度：O(log n)
//二分查找 -- 递归
func binarySearch(array []int, target int, startIndex int, endIndex int) int {
	//初始条件：二分发索引边界
	if startIndex > endIndex {
		return -1
	}
	//二分法指定的索引值
	mid := (startIndex + endIndex) / 2
	if array[mid] > target {
		return binarySearch(array, target, mid+1, endIndex)
	} else if array[mid] < target {
		return binarySearch(array, target, startIndex, mid-1)
	} else {
		return mid
	}
}

//二分查找 -- 迭代
func iterBinarySearch(array []int, target int, startIndex int, endIndex int) int {
	startIndex = 0
	endIndex = len(array) - 1
	//迭代条件
	for startIndex < endIndex {
		mid := (startIndex + endIndex) / 2
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	//默认返回值
	return -1
}
