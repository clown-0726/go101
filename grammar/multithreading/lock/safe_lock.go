package lock

import "sync"

type safeMap struct {
	safeMap map[int]int
	sync.Mutex
}

// 读数据的时候加锁
func (s *safeMap) Read(k int) (int, bool) {
	s.Lock()
	defer s.Unlock()

	result, ok := s.safeMap[k]
	return result, ok
}

// 写数据的时候加锁
func (s *safeMap) Write(k, v int) {
	s.Lock()
	defer s.Unlock()

	s.safeMap[k] = v
}

func SafeMapWrite() {
	s := safeMap{
		safeMap: map[int]int{},
		Mutex:   sync.Mutex{},
	}

	for i := 0; i < 100; i++ {
		go func() {
			s.Write(1, 1)
		}()
	}
}
