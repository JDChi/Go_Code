package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func TestRWMutex() {
	var counter Counter
	// more read
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println("read count: ", counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// less write
	for {
		counter.Incr()
		fmt.Println("write count")
		time.Sleep(time.Second)
	}
}
