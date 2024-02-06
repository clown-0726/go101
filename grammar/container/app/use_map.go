package app

import (
	"fmt"
)

func MapOps() {
	// 定义 Map
	var map1 map[string]string      // nil
	map2 := make(map[string]string) // empty map
	map3 := map[string]string{
		"a": "x",
		"b": "y",
		"c": "z",
	}
	fmt.Println(map1, map2, map3)

	// 遍历 map
	for k, v := range map3 {
		fmt.Println(k, v)
	}

	// 取值
	a := map3["a"]
	fmt.Println(a)

	// 如果从 map 中取不存在的 key 的时候，不会报错，而是返回对应类型的默认初始值
	// 同时会返回第二个 boolean 类型的参数来标识是否存在对应的 key
	if d, ok := map3["d"]; ok {
		fmt.Println(d)
	} else {
		fmt.Println("No key with d")
	}

	// 删除
	delete(map3, "c")

	// 除了 slice，map，function 的内建类型都可以作为 key
	// Struct 类型不包含上述字段，也可以作为 key
}
