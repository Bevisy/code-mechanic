package main

import (
	"fmt"
)

func main() {
	// 数组定义
	// 数组需要明确长度
	//var (
	//	a1 = [3]int{}
	//	a2 = [...]int{1, 2}
	//	b1 = [...]int{1, 2: 1, 3} // [1, 0, 1, 3]
	//)

	// 切片定义
	var (
		//	a []int // nil 切片,和nil相等，一般表示一个不存在的切片
		//	//b = []int{} // 空切片，和nil不相等，一般表示一个空的集合
		c = []int{1, 2, 3}
		//	//d = c[:2]
		e = c[0:len(c):cap(c)] // []前两个值是切片范围，最后一个值是切片容量
	//	//f = c[:0]
	//	//g = make([]int, 3)
	//	//h = make([]int, 2, 3)
	//	//i = make([]int, 0, 3)
	)
	//
	//fmt.Println(len(a), cap(a))
	fmt.Println(len(e), cap(e))

	// 切片修改
	//var a = []int{1, 2, 3}
	//fmt.Printf("%v\n", a)
	//a = append([]int{1}, a...) // a 需要解包
	//fmt.Printf("%v\n", a)
	//for i := range a {
	//	fmt.Printf("%d\n", a[i])
	//}

	// 切片中间插入元素
	// append 函数返回新的切片，所以支持链式操作
	//var a []int
	//a = append(a[:i], append([]int{x}, a[i:]...)...)     // 在第i个位置插入x
	//a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片

	// 利用copy 不创建临时切片的情况下，完成添加新元素的操作
	//a = append(a, 0)     // 切片扩展1个空间
	//copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	//a[i] = x             // 设置新添加的元素

	// 插入多个元素
	//a = append(a, x...)       // 为x切片扩展足够的空间
	//copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
	//copy(a[i:], x)            // 复制新添加的切片

}
