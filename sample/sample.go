package main

import "fmt"

func main() {
	type t interface{}
	var i t = 1
	fmt.Printf("%t\n", i.(int))
}
