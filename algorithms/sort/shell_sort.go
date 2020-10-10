package sort

//https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F
//希尔排序（Shellsort），也称递减增量排序算法，是插入排序的一种更高效的改进版本。希尔排序是非稳定排序算法
//算法时间复杂度O(n^2) -- 步长为 n/2
func shellSort(array []int) []int {
	for d := len(array) / 2; d > 0; d /= 2 {
		for i := d; i < len(array); i++ {
			for j := i; j >= d && array[j-d] > array[j]; j -= d {
				array[j], array[j-d] = array[j-d], array[j]
			}
		}
	}
	return array
}
