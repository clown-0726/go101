package app

import (
	"fmt"
	"sync"
	"time"
)

func MainDeferFinally() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	loopFunc()

	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// 必须把 defer 包裹在一个函数中，才能在函数执行完成之后执行释放
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc", i)
		}(i)
	}
}
