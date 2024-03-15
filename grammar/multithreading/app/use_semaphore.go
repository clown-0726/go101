package app

import (
	"fmt"
	"go101/grammar/multithreading/utils"
	"runtime"
	"sync"
	"time"
)

func MainSemaphore() {
	// 创建一个信号量，容量为 2
	sem := utils.NewSimpleSemaphore(2)

	// 创建一个 WaitGroup，用于等待所有协程完成
	var wg sync.WaitGroup

	// 开启 5 个协程
	for i := 1; i <= 500; i++ {
		wg.Add(1)
		//go worker(i, sem, &wg)
		go func(ind int) {
			defer wg.Done()

			// 获取信号量
			sem.Acquire()
			defer sem.Release()

			fmt.Println("worker", ind, "started", runtime.NumCPU(), runtime.NumGoroutine())
			time.Sleep(time.Second)
			fmt.Println("worker", ind, "finished")
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()
}
