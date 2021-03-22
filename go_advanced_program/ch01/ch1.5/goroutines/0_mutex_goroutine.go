package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		// 此处不加锁，部分情况下也会得到争取的值，因为实际争抢恰好没发生
		//total.Lock()
		total.value += i
		//total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	// 两个 worker 分别增加 total.value 的值，分别每个增加 5050
	for i := 0; i < 10; i++ {
		go worker(&wg) // 增加 worker数量，增大协程同时操作一个值的概率，成功触发资源争抢，得到结果为： 47341；多次运行，资源争抢发生
	}
	wg.Wait()

	fmt.Println(total.value)
}
