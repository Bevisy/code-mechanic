package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			c <- i
		}

		close(c)
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 6; i++ {
		fmt.Println(<-c)
	}
}
