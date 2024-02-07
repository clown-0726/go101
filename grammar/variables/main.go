package main

import (
	"fmt"
	"go101/grammar/variables/app"
)

func map_test() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 1
		},
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
}

func slice_test() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice %+v\n", mySlice)

	fullSlice := myArray[:]
	deletedSlice := deleteItem(fullSlice, 2)
	fmt.Printf("mySlice %+v\n", deletedSlice)
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1])
}

func slice_modify_test() {
	mySlice := []int{10, 20, 30, 40, 50}

	// 值传递，因此无法改变数组中的值
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

	// 通过下标的方式改变数组中的值
	for index, _ := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

}

func main() {
	// 变量定义
	app.UseVariableDefine()

	// 定义常量
	app.UseConst()
	app.UseEnums()

	// 使用 Flag
	app.UseFlag()

	//map_test()
	//slice_test()
	//slice_modify_test()
}
