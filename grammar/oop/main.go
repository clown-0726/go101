package main

import (
	"fmt"
	"go101/grammar/oop/app"
)

func main() {
	app.UseNodeMany()
	app.InOrderTraverse()

	fmt.Println("---------")

	// use q
	q := app.Queue{1, 2, 3}
	q.Push(4)
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	fmt.Println("---------")

	// PostOrder
	app.PostOrderTraverse()

	fmt.Println("---------")

	// PostOrderEmbedding
	app.PostOrderTraverseEmbedding()

	app.MainCalc()

	// Interface
	app.MainWhatIsInterface()

	app.MainInterfaceArray()
}

/*
类型扩充的方式：
定义别名
使用组合
*/
