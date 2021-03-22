package main

func main() {
	done := make(chan int)
	go func() {
		close(done) // 读取一个关闭的channel，会得到对应类型的默认值
	}()

	for i := 0; i < 10; i++ {
		println(<-done) // 输出 0
	}
}
