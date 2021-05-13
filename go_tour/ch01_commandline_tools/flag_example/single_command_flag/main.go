package main

import (
	"flag"
	"log"
)

// ./main -help
// flag 参数命令行配置模式，其中 "-" 等价于 "--"
// -flag：仅支持布尔类型。
// -flag x ：仅支持非布尔类型。
// -flag=x：均支持
func main() {
	var name string
	flag.StringVar(&name, "name", "test name", "help message")
	flag.StringVar(&name, "n", "test n", "help message")
	flag.Parse()

	log.Printf("name: %s\n", name)
}
