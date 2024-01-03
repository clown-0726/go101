package app

import "fmt"

func Main4ParamChange() {
	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)

	changeParameter(&para, "bbb")
	fmt.Println(para)

	cannotChangeParameter(para, "ccc")
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

// 传入指针，修改的是传入数据的值
func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

// go 语言只有值传递，para 是新的变量，因此原数据没有变化
func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}
