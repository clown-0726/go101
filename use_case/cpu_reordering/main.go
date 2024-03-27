package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func withoutCpuReordering() {
	index := 0
	for {
		index += 1

		var a, b int32 = 0, 0
		var x, y int32 = 0, 0

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			atomic.StoreInt32(&a, 1)
			atomic.StoreInt32(&x, atomic.LoadInt32(&b))
		}()

		go func() {
			defer wg.Done()

			atomic.StoreInt32(&b, 1)
			atomic.StoreInt32(&y, atomic.LoadInt32(&a))
		}()
		wg.Wait()

		if x == 0 && y == 0 {
			panic("CPU Reordering occurs!")
		} else {
			fmt.Println("Now processing in loop", index)
		}
	}
}

func withCpuReordering() {
	index := 0
	for {
		index += 1

		var a, b int32 = 0, 0
		var x, y int32 = 0, 0

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			a = 1
			x = b
		}()

		go func() {
			defer wg.Done()

			b = 1
			y = a
		}()
		wg.Wait()

		if x == 0 && y == 0 {
			panic("CPU Reordering occurs!")
		} else {
			fmt.Println("Now processing in loop", index)
		}
	}
}

func main() {
	// 在第 1738, 110002, 12987 次测试到了 CPU 指令重排序
	//runtime.GOMAXPROCS(1)
	withoutCpuReordering()

	// 一直到 1370370 也没测试到 CPU 指令重排序
	//withCpuReordering()
}
