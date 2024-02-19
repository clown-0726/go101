package main

import (
	"fmt"
	"go101/grammar/oop/app"
)

func main() {
	app.UseNodeMany()
	app.UseNode()

	fmt.Println("---------")

	// use q
	q := app.Queue{1, 2, 3}
	q.Push(4)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	fmt.Println("---------")

	// PostOrder
	app.PostOrder()

	fmt.Println("---------")

	// PostOrderEmbedding
	app.PostOrderEmbedding()

	app.MainCalc()
}

/*
类型扩充的方式：
定义别名
使用组合
*/
