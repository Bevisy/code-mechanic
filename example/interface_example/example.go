package main

import (
	"fmt"
	"sort"
)

type Array []int

func (arr Array) Len() int {
	return len(arr)
}

func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Rarray struct {
	Array
}

func (arr Rarray) Less(i, j int) bool {
	return arr.Array[j] < arr.Array[i]
}

func main() {
	arr := Array{1, 3, 4, 2}
	sort.Sort(arr) // 升序
	fmt.Printf("%v\n", arr)

	// rarr := sort.Reverse(arr)
	// sort.Sort(rarr) // 降序
	// fmt.Printf("%v\n", rarr)

	rarr := Rarray{
		Array: arr,
	}
	sort.Sort(rarr)
	fmt.Printf("%v\n", rarr)
}
