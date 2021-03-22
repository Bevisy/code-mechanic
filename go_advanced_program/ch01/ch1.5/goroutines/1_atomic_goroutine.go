package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用互斥锁保护数值型的共享资源，麻烦且效率低下
// sync/atomic 包对原子操作提供丰富的支持

var total1 int64

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i int64
	for i = 0; i <= 100; i++ {
		atomic.AddInt64(&total1, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker1(&wg)
	}

	wg.Wait()

	fmt.Println(total1)
}
