package main

import "fmt"

type animal struct {
	sound string
}

// cat 结构体包含 animal 结构体内容
type cat struct {
	action string
	animal
}

func main() {
	var c = new(cat)
	c.animal.sound = "meow"
	c.action = "lay"
	c.sound = "meow2" // 最后赋值，覆盖之前的赋值

	fmt.Println(c.sound)
}
