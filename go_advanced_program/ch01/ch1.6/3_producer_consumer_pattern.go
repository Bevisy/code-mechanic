package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 生产者和消费者模型
func producer(num int, out chan<- int) {
	for i := 0; ; i++ {
		out <- num * i
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	go producer(3, ch)
	go producer(5, ch)

	go consumer(ch)

	// ctrl + C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
