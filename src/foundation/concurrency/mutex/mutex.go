package mutex

import (
	"fmt"
	"sync"
)

func WithoutMutex() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println("WithoutMutex counter count: ", count)

}

type Counter struct {
	sync.Mutex
	Count uint64
}

type SafeCounter struct {
	mu    sync.Mutex
	count uint64
}

func (counter *SafeCounter) Incr() {
	counter.mu.Lock()
	counter.count++
	counter.mu.Unlock()
}

func (counter *SafeCounter) Count() uint64 {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.count
}

func FixWithMutex() {
	var counter Counter
	var safeCounter SafeCounter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()

				safeCounter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println("FixWithMutex counter count: ", counter.Count)
	fmt.Println("FixWithMutex safeCounter count: ", safeCounter.count)
}
