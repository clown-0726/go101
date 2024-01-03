package pkgs

func SingleChannel() {
	var c = make(chan int)
	go prod(c)
	go customer(c)
}

func prod(ch chan<- int) {
	for {
		ch <- 1
	}
}

func customer(ch <-chan int) {
	for {
		<-ch
	}
}
