package main

import (
	"flag"
	"log"
)

var name string

// 一级命令：
// go run main.go --help
// 子命令模式：
// go run main.go go --help
// go run main.go php --help
func main() {
	flag.StringVar(&name, "name", "main flag", "level_1 flag")
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go language", "help message")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "PHP language", "help message")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s\n", name)
}
