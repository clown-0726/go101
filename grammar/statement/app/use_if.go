package app

import (
	"fmt"
	"io/ioutil"
)

func UseIf() {
	filename := "grammar/statement/app/abc.txt"

	// if 条件里可以赋值
	// if 条件里赋值变量的作用域就是当前的 if 中有效
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
}
