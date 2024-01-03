package pkgs

import (
	"fmt"
	"time"
)

func CustomerProducer() {
	messages := make(chan int, 10)
	done := make(chan bool)

	// defer close(done)

	// customer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("Child process interrupt...")
				return
			default:
				fmt.Println("Send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit...")
}
