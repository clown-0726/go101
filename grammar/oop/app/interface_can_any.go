package app

import "fmt"

// interface{} 表示任何类型

type QueueAny []interface{}

func (q *QueueAny) Push(v interface{}) {
	// q 所指向的 object 本身变掉了
	*q = append(*q, v.(int))
}

func (q *QueueAny) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *QueueAny) IsEmpty() bool {
	return len(*q) == 0
}

func MainUseQueueAny() {
	q := QueueAny{1, 2, 3}
	// q.Push("4") // 可以传递任何类型，但是运行会出错
	q.Push(4)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
}
