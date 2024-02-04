package app

import "fmt"

func UseArray() {
	var arr1 [5]int                 // 直接定义数组，初始值都为默认值 0
	arr2 := [3]int{1, 2, 3}         // 定义明确个数的数组，并赋初始值
	arr3 := [...]int{1, 2, 3, 4, 5} // 定义数组，并赋初始值，由系统判断数组个数
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func UseArrayRange(arr [5]int) {
	// 通过 range 遍历
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// 通过 for 遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
