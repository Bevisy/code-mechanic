package main

import "fmt"

// 指针间复制传递的是地址
func main() {
	var a = [...]int{1, 2, 3}
	var b = &a // b 是指向 a 的指针

	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])

	for _, v := range b {
		fmt.Println(v)
	}
}
