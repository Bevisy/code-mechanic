package main

import (
	"fmt"
	"os"
	"strings"
)

//type fmt.Stringer interface {
//	String() string
//}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}
