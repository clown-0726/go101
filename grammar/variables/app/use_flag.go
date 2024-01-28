package app

import (
	"flag"
	"fmt"
)

/*
	flag 包提供了一种解析命令行参数的方式，包支持以下类型的命令行标志参数：
	- 布尔型标志参数
	- 字符串型标志参数
	- 整型标志参数
*/

func UseFlag() {
	// name 是标志参数，后面的使用都要通过 name 的指针使用
	name := flag.String("name", "World", "Specify the name you want to say?")

	fmt.Println("Input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
