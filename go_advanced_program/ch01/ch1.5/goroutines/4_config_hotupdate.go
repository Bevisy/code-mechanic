package main

import (
	"sync/atomic"
	"time"
)

// 一个简化的生产者消费者模型：后台线程生成最新的配置信息；
// 前台多个工作者线程获取最新的配置信息。所有线程共享配置信息资源

func loadConfiguration() interface{} {
	return new(interface{})
}

func main() {
	var config atomic.Value // 保存当前配置

	config.Store(loadConfiguration()) // 初始化配置信息

	go func() { // 启动后台协程，加载更新后的配置信息
		time.Sleep(time.Second * 1)
		config.Store(loadConfiguration())
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load() // 用于处理请求的工作者线程始终采用最新的配置信息
				//...
			}
		}()
	}
}
