package main

import (
	"sync"
)

var (
	instance1 *singleton
	once      sync.Once
)

func Instance1() *singleton {
	once.Do(func() { // 单例模式，只有初次调用会新建对象 instance1 ，否则均返回以新建的 instance1
		instance1 = &singleton{}
	})
	return instance
}
