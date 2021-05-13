package main

import (
	"fmt"
)

// 动态类型逃逸
//
// 函数参数类型为interface，编译期间难以确认其参数的具体类型，也会产生逃逸
func main() {
	fmt.Println("hello") // ./main.go:6:14: "hello" escapes to heap
}
