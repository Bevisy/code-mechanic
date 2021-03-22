package main

//#include <stdio.h> // 包含 c语言头文件<stdio.h>
import (
	"C"
)

func main() {
	C.puts(C.CString("hello, world\n"))
}
