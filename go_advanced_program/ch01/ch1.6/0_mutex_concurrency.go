package main

import (
	"fmt"
	"sync"
)

func main() {
	// 通过两次加锁，当第二次加锁时，需要等待锁释放
	var m sync.Mutex

	m.Lock()
	go func() {
		fmt.Println("hello")
		m.Unlock()
	}()

	m.Lock()
}
