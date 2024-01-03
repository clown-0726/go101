package lock

import (
	"fmt"
	"sync"
)

func WaitGroupMock() {
	// 通过 channel 完成线程阻塞
	c := make(chan bool, 100)

	// 往 channel 中加入数据
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	// 在 channel 中如果没有数据拿，则阻塞
	for i := 0; i < 100; i++ {
		<-c
	}
}

func WaitGroup() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			// 每当一个线程执行完成，执行一次 Done 操作
			wg.Done()
		}(i)
	}
	// 阻塞，当达到定义的 WaitGroup 上限之后继续执行
	wg.Wait()
	fmt.Println("Main process go...")
}
