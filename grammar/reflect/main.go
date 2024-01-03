package main

import (
	"fmt"
	"reflect"
)

func toReflect() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myMap["c"] = "d"

	t := reflect.TypeOf(myMap)
	fmt.Println(t)

	v := reflect.ValueOf(myMap)
	fmt.Println(v)
}

func toReflect2() {
	myStruct := T{A: "abc"}

	// 通过反射可以得到所有的属性
	v1 := reflect.ValueOf(myStruct)
	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v1.Field(i))
	}

	// 通过反射可以得到所有的方法
	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}

	// 也可以进行方法的调用
	result := v1.Method(0).Call(nil)
	fmt.Println(result)
}

type T struct {
	A string
}

func (t T) String() string {
	return t.A + "1"
}

func main() {
	toReflect()
	toReflect2()
}
