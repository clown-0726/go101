package main

import (
	"go101/grammar/multithreading/lock"
	"time"
)

func main() {
	// ------- thread & channel
	//pkgs.LoopFunc()
	//pkgs.ChannelRun()
	//pkgs.ChannelClose()
	//pkgs.ChannelStop()
	//pkgs.ChannelSelect()
	//pkgs.CustomerProducer()
	//pkgs.SingleChannel()
	//pkgs.CtxTest()

	// ------- lock
	//lock.SafeMapWrite()
	//lock.WaitGroupMock()
	//lock.WaitGroup()
	lock.InvokeCond()

	time.Sleep(time.Second)
}
