package main

import (
	"fmt"
	"go101/grammar/oom/app"
)

func main() {
	app.UseNodeMany()
	app.UseNode()

	// use q
	q := app.Queue{1, 2, 3}
	q.Push(4)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	// PostOrder
	app.PostOrder()
}

/*
类型扩充的方式：
定义别名
使用组合
*/
