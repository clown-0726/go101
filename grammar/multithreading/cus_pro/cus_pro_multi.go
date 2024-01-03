package cus_pro

import (
	"fmt"
	"sync"
)

// 生产者
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		ch <- i
		fmt.Printf("Producer has produced: %d\n", i)
	}
	close(ch)
}

// 消费者
func consumer(id string, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Printf("Consumer %s has consumed: %d\n", id, i)
	}
}

func Main4CusProMulti() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4) // 一个生产者和两个消费者，所以总共是3个goroutine
	go producer(ch, &wg)
	go consumer("A", ch, &wg)
	go consumer("B", ch, &wg)
	go consumer("C", ch, &wg)
	wg.Wait()
}
