package main

import (
	"fmt"
)

// Talk 接口
type talk interface {
	Hello() string
	Talk(heard string) (answer string, err error)
}

// chatBot 接口包含Talk接口
// chatBot接口包含Talk接口全部方法，声明形式如下
type chatBot interface {
	talk // 包含接口，即包含接口全部方法
	Name() string
	Report() string
}

// 实现 Talk 接口，同时实现 chatBot 接口
type myTalk string

func (t *myTalk) Hello() string {
	return "hello"
}

func (t *myTalk) Talk(heard string) (answer string, err error) {
	return "hello " + string(*t), nil
}

func (t *myTalk) Name() string {
	return string(*t)
}

func (t *myTalk) Report() string {
	return ""
}

func main() {
	name := "cherry"
	t := myTalk(name)
	fmt.Println(t.Name())
}
