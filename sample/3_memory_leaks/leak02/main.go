package main

// 栈空间不足逃逸
//
// 当切片长度扩大到10000时就会逃逸。实际上当栈空间不足以存放当前对象时或无法判断当前切片长度时会将对象分配到堆中
func main() {
	s1 := make([]int, 1000) // make([]int, 1000) does not escape
	for index, _ := range s1 {
		s1[index] = index
	}

	s2 := make([]int, 10000) // make([]int, 10000) escapes to heap
	for index, _ := range s2 {
		s2[index] = index
	}
}
