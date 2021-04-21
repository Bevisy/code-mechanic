package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	initChan := make(chan int, 1)
	t := time.NewTicker(time.Second)
	go func() {
		for _ = range t.C {
			select {
			case initChan <- 1:
			case initChan <- 2:
			case initChan <- 3:
			}
		}
	}()

	var sum int
	for i := range initChan {
		if sum < 10 {
			fmt.Println(i)
			sum += i
		} else {
			fmt.Printf("%d\n", sum)
			os.Exit(0)
		}
	}
}
