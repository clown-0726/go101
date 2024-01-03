package pkgs

import (
	"fmt"
	"math/rand"
	"time"
)

// 遍历通道缓冲区
func ChannelRun() {
	// 声明一个 Channel
	ch := make(chan int, 10)

	// 在这个新的线程中放入数据
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("Putting: ", n)
			ch <- n
		}
		close(ch)
	}()

	fmt.Println("Hello from ChannelRun!")

	// 主线程中进行数据消费
	for v := range ch {
		fmt.Println("Receiving: ", v)
	}
}
