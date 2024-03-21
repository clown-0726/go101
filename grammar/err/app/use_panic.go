package app

import (
	"errors"
	"fmt"
	"os"
)

func MainPanicSimple() {
	fmt.Println("Some business code are running with err!")
	panic("Invoke panic on propose!") // panic 会退出执行程序
}

func MainErrWithType() {
	// 尝试以互斥的方式打开文件，如果文件已经存在，则报错
	file, err := os.OpenFile("abc.txt", os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())

		// 进行特定的错误捕获
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			// 兜底的错误捕获
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()
}

func throwErr() error {
	return errors.New("I a error")
}

func MainThrowCatch() {
	err := throwErr()
	if err != nil {
		panic(err)
	}
}
