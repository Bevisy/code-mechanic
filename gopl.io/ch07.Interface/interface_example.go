package main

import (
	"fmt"
	"sort"
)

// 二维整型数组，第二数组元素末尾大小排序
type dataType [][]int

func (p dataType) Len() int {
	return len(p)
}

func (p dataType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p dataType) Less(i, j int) bool {
	return p[i][len(p[i])-1] < p[j][len(p[j])-1]
}

func main() {
	data := dataType{{1, 3}, {1, 2}}
	sort.Sort(data)
	fmt.Println(data)
}
