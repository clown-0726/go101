package main

import (
	"flag"
	"fmt"
	"os"
)

// 每个 Go 语言程序都应该有个 main package
// Main package 里的 main 函数是 Go 语言程序入口

func main() {
	// 通过 flag 进行参数解析
	name := flag.String("name", "world", "specify the name to greet")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)

	err, result := DuplicateString("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

// 函数可以有多个返回值
func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}
