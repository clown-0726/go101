package app

import "fmt"

// 通过定义别名来扩充类型

type Queue []int

func (q *Queue) Push(v int) {
	// q 所指向的 object 本身变掉了
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func MainUseQueue() {
	q := Queue{1, 2, 3}
	q.Push(4)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
}
