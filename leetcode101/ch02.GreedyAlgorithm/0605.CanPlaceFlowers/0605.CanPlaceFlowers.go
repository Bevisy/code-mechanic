package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	if len(flowerbed) == 0 {
		return false
	} else if len(flowerbed) == 1 {
		if flowerbed[0] == 0 && n == 1 {
			count++
		} else if flowerbed[0] == 0 && n == 0 {
			count++
		}
	} else {
		for i := 0; i < len(flowerbed); i++ {
			// 数组首尾特殊处理
			if i == 0 {
				if flowerbed[0] == 0 && flowerbed[1] == 0 {
					count++
					flowerbed[i] = 1
				}
			} else if i == len(flowerbed)-1 {
				if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
					count++
					flowerbed[i] = 1
				}
			} else { // 处理数组中间位置
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
					count++
					flowerbed[i] = 1 // 统计可种花的地，置为种花，避免后续误判
				}
			}
		}
	}

	return count == n
}
