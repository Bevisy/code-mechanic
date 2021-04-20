package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d\n", n)
		}(i * 2) // defer 被执行时，被求值
	}
}
