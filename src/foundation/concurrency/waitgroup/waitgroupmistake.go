package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func MistakeAdd() {
	var wg sync.WaitGroup
	wg.Add(10)
	wg.Add(-10)
	wg.Add(-1)
}

func MistakeDone() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Done()
}

func MistakeAddTiming() {
	var wg sync.WaitGroup
	go doSomething(100, &wg)
	go doSomething(110, &wg)
	go doSomething(120, &wg)
	go doSomething(130, &wg)

	wg.Wait()
	fmt.Println("Done")

}

func doSomething(millsecs time.Duration, wg *sync.WaitGroup) {
	duration := millsecs * time.Millisecond
	time.Sleep(duration)
	wg.Add(1)
	fmt.Println("go running, duration: ", duration)
	wg.Done()
}

func MistakeReuseWaitGroupBeforeDone() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
