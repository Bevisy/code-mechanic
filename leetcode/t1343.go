package leetcode

func numOfSubarrays(arr []int, k int, threshold int) int {
	var count int // 计数标识
	if len(arr) < k {
		return 0
	} else {
		for i := 0; i <= len(arr)-k; i++ {
			if ave(arr[i:i+k]) >= threshold {
				count++
			} else {
				continue
			}
		}
	}
	return count
}

// 数组求和
func sum(arr []int) int {
	var sums int
	for _, v := range arr {
		sums += v
	}
	return sums
}

// 数组均值
func ave(arr []int) int {
	var average int
	average = sum(arr) / len(arr)
	return average
}
