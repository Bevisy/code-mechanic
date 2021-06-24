package main

import "fmt"

func main() {
	var i uint64 = 1 << 13
	fmt.Println(sovRaft(i))
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
