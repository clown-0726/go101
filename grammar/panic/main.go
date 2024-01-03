package main

import "fmt"

func panicTest() {
	defer func() {
		fmt.Println("defer func is called!")
		if err := recover(); err != nil { // 相当于捕获到异常
			fmt.Println(err)
		}
	}()

	fmt.Println("Other business code running here!")

	panic("a panic is triggered!") // 相当于进行异常抛出
}

func main() {
	panicTest()
	/**
	1. Other business code running here!
	2. defer func is called!
	3. a panic is triggered!
	*/
}
