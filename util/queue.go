package util

type Queue struct {
	Queue []int
}

// 入队
func (q *Queue) Append(target int) {
	q.Queue = append(q.Queue, target)
}

// 出队
func (q *Queue) Pop() int {
	p := q.Queue[0]
	q.Queue = q.Queue[1:]
	return p
}

// 队空检查
func (q *Queue) IsNil() bool {
	if len(q.Queue) <= 0 {
		return true
	} else {
		return false
	}
}
