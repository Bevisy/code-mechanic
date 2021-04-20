package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 不断读取输入，匹配并输出对应策略
	fmt.Println("Please input your name:")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Hello %s, what can do for you?\n", name[:len(name)-1])

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("read input: %s\n", err)
			continue
		}

		input = input[:len(input)-1]
		switch input {
		case "bye", "nothing":
			fmt.Println("Bye!")
			break
		case "exit":
			fmt.Println("exiting...")
			os.Exit(0)
		case "":
			continue
		default:
			fmt.Println("Sorry, I couldn't catch you!")
		}
	}
}
