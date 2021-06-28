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
			} else {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
					count++
					flowerbed[i] = 1
				}
			}
		}
	}

	return count == n
}
