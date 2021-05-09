package goroutine

import "sync"

func basicAppend(b []byte) {
	s := make([]byte, 0, len(b))
	for _, v := range b {
		s = append(s, v)
	}
}

func useGoroutineAppend(b []byte) {
	s := make([]byte, 0, len(b))
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, v := range b {
		wg.Add(1)
		go func(v byte) {
			mu.Lock()
			s = append(s, v)
			defer mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
}