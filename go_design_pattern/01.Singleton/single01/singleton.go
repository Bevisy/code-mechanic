package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

// 全局实例
var singleInstance *single

// 获取实例方法
func getInstance() *single {
	// 最开始时会有 nil检查，确保 single­Instance单例实例在最开始时为空。这是为了防止在每次调用 get­Instance方法时都去执行消耗巨大的锁定操作。 如果检查不通过，则就意味着 single­Instance字段已被填充。
	if singleInstance == nil {
		lock.Lock() // single­Instance结构体将在锁定期间创建
		defer lock.Unlock()
		// 在获取到锁后还会有另一个 nil 检查。这是为了确保即便是有多个协程绕过了第一次检查，也只能有一个可以创建单例实例。否则，所有协程都会创建自己的单例结构体实例
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{} // 实例初始化
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			getInstance()
			wg.Done()
		}()
	}

	wg.Wait()
}
