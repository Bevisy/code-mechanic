package sort

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	var a = mergeSort(arr[:middle])
	var b = mergeSort(arr[middle:])
	return merge(a, b)
}

func merge(arr1 []int, arr2 []int) []int {
	var result = make([]int, len(arr1)+len(arr2))
	var i = 0
	var j = 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[i+j] = arr1[i]
			i++
		} else {
			result[i+j] = arr2[j]
			j++
		}
	}

	for i < len(arr1) {
		result[i+j] = arr1[i]
		i++
	}

	for j < len(arr2) {
		result[i+j] = arr2[j]
		j++
	}

	return result
}
