package pkgs

import "fmt"

func ChannelSelect() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		// 一直循环从不同 channel 中读取数据
		for {
			select {
			case v := <-chan1:
				println("I am case1 chan1", v)
			case v := <-chan2:
				println("I am case2 chan2", v)
			default:
				println("I am case default")
				return
			}
		}
	}()

	// 主线程放置数据
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		chan1 <- i
	}
	fmt.Println("To close chan1")
	close(chan1)
}
