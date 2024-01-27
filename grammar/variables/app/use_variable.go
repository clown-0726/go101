package app

import "fmt"

func UseVariableDefine() {
	// 先写变量名后写类型对于复合变量的定义会有一定优势
	var name string = "Xiaoming"
	var age int = 18
	var address string
	// %s 代表字符串类型，%q 表示 quote，%d 表示数字类型
	fmt.Printf("%s %q %d\n", name, address, age)

	// 快捷定义变量，可以同时定义不同的类型
	a, b, c := "abc", 3, true
	fmt.Println(a, b, c)

	// 可以通过 () 快捷定义多个变量
	var (
		x      = "xyz"
		y      = 9
		z bool = false
	)
	fmt.Println(x, y, z)
}
