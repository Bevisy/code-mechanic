package utils

import "math/rand"

func arrayGenerator(length int) []int {
	var result []int
	for i := 0; i < length; i++ {
		result = append(result, rand.Intn(100))
	}
	return result
}
