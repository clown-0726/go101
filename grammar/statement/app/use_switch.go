package app

import "fmt"

func foo() {
	fmt.Println("I am foo")
}

func UseSwitch() {
	// switch 会自动 break，否则使用 fallthrough
	var1 := "b"
	switch var1 {
	case "a": // 空分支
	case "b":
		fallthrough // 执行 case3 中的 foo()
	case "c":
		foo()
	default: // 默认分支
		panic("err...")
	}

}
