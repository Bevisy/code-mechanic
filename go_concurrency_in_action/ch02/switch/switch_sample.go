package main

import (
	"fmt"
	"strings"
)

// 支持单个值
func switch1(content string) string {
	switch content {
	default:
		return "UNKNOWN TYPE"
	case "Python":
		return "python"
	case "Go":
		return "go"
	}
}

// 支持值初始化
func switch2(content string) string {
	switch lang := strings.TrimSpace(content); lang {
	default:
		return "UNKNOWN TYPE"
	case "Python":
		return "python"
	case "Go":
		return "go"
	}
}

// fallthrough 想下一个case语句转移流程控制权
func switch3(content string) string {
	switch content = strings.TrimSpace(content); content {
	default:
		return "UNKONWN TYPE"
	case "Java":
		fallthrough
	case "Python":
		return "python"
	case "Go":
		return "go"
	}
}

// case 表达式支持存在多个
func switch4(content string) string {
	switch content = strings.TrimSpace(content); content {
	default:
		return "UNKONWN TYPE"
	case "Java":
		fallthrough
	case "Python":
		return "python"
	case "Go", "GO", "Golang", "golang":
		return "go"
	}
}

// switch 支持使用 break 退出当前语句
func switch5(content string) string {
	switch content = strings.TrimSpace(content); content {
	default:
		return "UNKONWN TYPE"
	case "Java":
		fallthrough
	case "Python":
		break // 跳出 switch 代码块
	case "Go", "GO", "Golang", "golang":
		return "go"
	}

	return ""
}

// switch 可以用于接口类型判断，而不是用于值
// fallthrough 不允许出现在类型switch语句中
func switch6(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case int:
		return "int"
	default:
		return "Unkonwn"
	}
}
func main() {
	var s = "  Python  "
	fmt.Println(switch6(s))
}
