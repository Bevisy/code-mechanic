package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputReade := bufio.NewReader(os.Stdin)
	fmt.Println("Input your name:")
	input, err := inputReade.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// 对input进行切片，去除末尾字符
	fmt.Printf("Hello %s\n", input[:len(input)-1])
}
