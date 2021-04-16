package main

import (
	"fmt"
	"reflect"
)

func main() {
	//t := reflect.TypeOf(1i)
	//fmt.Println(t.String())
	//fmt.Println(t)

	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
}
