package app

import (
	"fmt"
	"reflect"
	"runtime"
)

func MainDefOfFunc() {
	// 返回值类型写在最后面
	basicFun("abc")

	// 函数可以有多个返回值
	err, result := duplicateString("abc")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	// 函数的返回值可以指定名称
	res, remainder := divAB(9, 4)
	fmt.Println(res, remainder)

	// 函数作为参数
	res2 := apply(func(i1 int, i2 int) int {
		return i1 + i2
	}, 3, 4)
	fmt.Println(res2)

	// 没有默认参数，可选参数
	// 但是有可变参数列表
	res3 := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(res3)
}

// 返回值类型写在最后面
func basicFun(text string) string {
	return text
}

// 函数可以有多个返回值
func duplicateString(input string) (error, string) {
	if input == "abc" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input
}

// 函数的返回值可以指定名称
func divAB(a, b int) (result int, remainder int) {
	result = a / b
	remainder = a % b
	return
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	// 通过反射得到调用函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function is %s\n", opName)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
