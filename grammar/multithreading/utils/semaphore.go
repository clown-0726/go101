package utils

type SimpleSemaphore chan struct{}

func NewSimpleSemaphore(n int) SimpleSemaphore {
	return make(SimpleSemaphore, n)
}

func (ss *SimpleSemaphore) Acquire() {
	*ss <- struct{}{}
}

func (ss *SimpleSemaphore) Release() {
	<-*ss
}
