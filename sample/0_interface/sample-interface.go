package main

import "fmt"

type Animal interface {
	Quark() string // 函数名、输入参数、返回值 需保持一致，才算是实现接口
	Eat()
}

type Cat struct {
	sound string
}

func (cat *Cat) Quark() string {
	return cat.sound
}

func (cat *Cat) Eat() {
	fmt.Print("fish~~~")
}

func main() {
	var cat Animal = &Cat{"meow"} // 使用接口声明类型；接口实现为指针类型，接收参数为指针
	cat.Eat()
}
