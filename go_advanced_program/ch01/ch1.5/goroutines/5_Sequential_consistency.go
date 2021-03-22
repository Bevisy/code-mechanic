package main

func init() { // init()是一种特殊函数，可在同一文件内被多次创建
 println("hello")
}

func init() {
	print("HELLO")
}

func main() {
	done := make(chan int)
	//var done chan int

	go func() {
		println("hello")
		done <- 1
	}()

	<- done
	//select {}
}
