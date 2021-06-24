package main

import "fmt"

func main() {
	fmt.Printf("%d\n", n += 1 + sovRecord(uint64(m.Index)))
}

func sovRecord(x uint64) (n int) {
	for {
		n++
		// fmt.Printf("n: %d\n", n)
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
