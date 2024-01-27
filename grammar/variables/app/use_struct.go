package app

// 定义结构体
type MyType struct {
	Name string
}

// 传入结构体指针
func printMyType(t *MyType) {
	println(t.Name)
}

func UseStruct() {
	// 使用结构体
	t := MyType{Name: "test"}
	printMyType(&t)
}
