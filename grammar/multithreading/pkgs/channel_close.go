package pkgs

import "fmt"

func ChannelClose() {
	ch := make(chan int)
	defer close(ch)

	// 可以合并到下面一起写 v, notClosed := <-ch
	if v, notClosed := <-ch; notClosed {
		fmt.Println(v)
	}
}

func ChannelStop() {
	done := make(chan bool)
	go func() {
		for {
			fmt.Println("Enter the for loop.")

			select {
			case <-done:
				fmt.Println("Done is triggerred, exit child go routine!")
				return
			}
		}
	}()
	close(done)
}
