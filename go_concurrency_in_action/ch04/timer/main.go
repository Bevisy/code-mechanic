package main

import (
	"fmt"
	"time"
)

func main() {
	initChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
	}()

	select {
	case <-initChan:
		fmt.Println("success")
	case <-time.NewTimer(time.Millisecond * 500).C:
		fmt.Println("timeout.")
	}
}
