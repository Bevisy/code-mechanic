package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) // 局部变量逃逸到堆
	s.Age = age
	s.Name = name

	return s
}

// 如果可能，Golang 编译器会将函数的局部变量分配到函数栈帧（stack frame）上。 然而，如果编译器不能确保变量在函数 return之后不再被引用，编译器就会将变量分配到堆上。而且，如果一个局部变量非常大，那么它也应该被分配到堆上而不是栈上。
// https://zhuanlan.zhihu.com/p/113643434
//
// go build -gcflags="-m" mai.go
// ./main.go:11:10: new(Student) escapes to heap 代表该行内存分配发生了逃逸 栈 --> 堆
func main() {
	s := StudentRegister("John", 18)

	fmt.Printf("name: %s, age: %d\n", s.Name, s.Age)
}
