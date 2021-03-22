package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 { // 如果 initialized 为1 ，则直接返回已有的 instance
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{} // 如果 initialized 不为1 ，则新建instance，否则直接返回已有的 instance
	}
	return instance
}
