package leetcode

// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i] // 计算所需的另外一个目标数字
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		} // 目标数字存在则返回下标，否则将数字作为 key，原数组下标作为 value 记录在map m 中
		m[nums[i]] = i
	}
	return nil
}
