package main

import "go101/grammar/functions/app"

// 每个 Go 语言程序都应该有个 main package
// Main package 里的 main 函数是 Go 语言程序入口

func main() {
	// 函数的基础语法
	//app.MainDefOfFunc()

	// 函数式编程
	app.MainInvokeAdder() // 累加器
	app.MainInvokeAdder2()
	app.MainFibo()
}
