package main

import "os"
import "fmt"

func getPid() int {
	pid := os.Getpid()
	//ppid := os.Getppid()
	return pid
}

func main() {
	fmt.Println(getPid())
}
