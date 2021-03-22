package main

import (
	"fmt"
)

func main() {
	var a = []interface{}{123, "abc"}
	Print(a)    // 不解包 [123 abc]
	Print(a...) // 解包 123 abc
}

// 当可变参数是一个空接口类型，传入参数是否解包可变参数将导致不同的结果
func Print(a ...interface{}) {
	fmt.Println(a...)
}
