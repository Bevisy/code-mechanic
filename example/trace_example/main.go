package main

import (
	"os"
	"runtime/trace"
)

// go run main.go 2> trace.out
// go tool trace trace.out
//
//
//View trace：查看跟踪
//Goroutine analysis：Goroutine 分析
//Network blocking profile：网络阻塞概况
//Synchronization blocking profile：同步阻塞概况
//Syscall blocking profile：系统调用阻塞概况
//Scheduler latency profile：调度延迟概况
//User defined tasks：用户自定义任务
//User defined regions：用户自定义区域
//Minimum mutator utilization：最低 Mutator 利用率
//
// https://mp.weixin.qq.com/s/7DY0hDwidgx0zezP1ml3Ig
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "Go travel."
	}()

	<-ch
}
