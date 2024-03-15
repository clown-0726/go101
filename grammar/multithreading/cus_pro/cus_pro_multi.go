package cus_pro

import (
	"fmt"
	"sync"
	"time"
)

// 生产者
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("Producer has produced: %d\n", i)
	}

	close(ch)
}

// 消费者
func consumer(id string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("I am done...", id)
			return
		}

		// Start deal with data
		time.Sleep(time.Second * 3)
		fmt.Printf("Consumer %s has consumed: %d\n", id, data)
	}
}

func Main4CusProMulti() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4) // 一个生产者和多个消费者
	go producer(ch, &wg)

	go consumer("A", ch, &wg)
	go consumer("B", ch, &wg)
	go consumer("C", ch, &wg)

	wg.Wait()
}
