package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			fmt.Println("hello")
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
