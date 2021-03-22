package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 过滤能被素数整除的数
// 每次执行，会生成新的管道
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

//先是调用GenerateNatural()生成最原始的从2开始的自然数序列(第一个管道)。
//然后开始一个100次迭代的循环，希望生成100个素数。
//在每次循环迭代开始的时候，管道中的第一个数必定是素数，我们先读取并打印这个素数。
//然后基于管道中剩余的数列，并以当前取出的素数为筛子过滤后面的素数。
//不同的素数筛子对应的管道是串联在一起的。

//
func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 1000; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}
