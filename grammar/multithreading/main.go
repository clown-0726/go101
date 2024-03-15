package main

import (
	"go101/grammar/multithreading/cus_pro"
)

func main() {
	// ------- thread & channel
	//pkgs.LoopFunc()
	//pkgs.ChannelRun()
	//pkgs.ChannelClose()
	//pkgs.ChannelStop()
	//pkgs.ChannelSelect()
	//pkgs.SingleChannel()
	//pkgs.CtxTest()

	// ------- lock
	//lock.SafeMapWrite()
	//lock.WaitGroupMock()
	//lock.WaitGroup()
	//lock.InvokeCond()

	cus_pro.Main4CusProMulti()
	//cus_pro.CustomerProducer()

	//app.MainSemaphore()
}
