package main

import (
	"fmt"
	"go101/grammar/container/app"
)

func main() {
	fmt.Println("Container...")
	// 数组都是值传递类型，因此比如调用 func f(arr [10]int) 会拷贝数组
	// 数组定义
	//app.UseArray()
	// 数组遍历
	//app.UseArrayRange([5]int{1, 2, 3, 4, 5})

	// Slice 基础
	//app.SliceBasic()
	//app.SliceUpdate()
	//app.ReSliceExtendingSlice()
	app.SliceOps()

	// Map 基础
	//app.MapOps()
}
