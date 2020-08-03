package util

type Stack struct {
	Stack []int // 切片
}

// 入栈
func (stack *Stack) Append(target int) {
	stack.Stack = append(stack.Stack, target)
}

// 出栈
func (stack *Stack) Pop() int {
	p := stack.Stack[len(stack.Stack)-1]
	stack.Stack = stack.Stack[:len(stack.Stack)-1]
	return p
}

// 栈深度
func (stack *Stack) Deepth() int {
	return len(stack.Stack)
}

// 空栈检查
func (stack *Stack) IsNil() bool {
	if len(stack.Stack) <= 0 {
		return true
	} else {
		return false
	}
}
