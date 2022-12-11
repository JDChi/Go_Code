package mutex

import (
	"fmt"
	"sync"
	"time"
)

func MistakeNoLockUnlockPairing() {
	// cause panic
	// we can't call Unlock before Lock
	var mu sync.Mutex
	mu.Unlock()
	fmt.Println("Hello World")
}

// MistakeCopyUsedMutex
// fatal error: all goroutines are asleep - deadlock!
func MistakeCopyUsedMutex() {
	var count int
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	count++
	func(_mu sync.Mutex) {
		_mu.Lock()
		defer _mu.Unlock()
		fmt.Println("MistakeCopyUsedMutex in foo")
	}(mu)
}

// MistakeReentrantLock
// In java, if a thread have got a lock, and it wants to get again, it will lock successfully.
// But in Go, there is no ReentrantLock
func MistakeReentrantLock() {
	l := &sync.Mutex{}

	// fatal error: all goroutines are asleep - deadlock!
	func(_l sync.Locker) {
		fmt.Println("MistakeReentrantLock in foo")
		l.Lock()
		func(_p sync.Locker) {
			_p.Lock()
			fmt.Println("MistakeReentrantLock in bar")
			_p.Unlock()
		}(_l)
	}(l)
}

func MistakeDeadLock() {
	var psCertificate sync.Mutex
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		psCertificate.Lock()
		defer psCertificate.Unlock()

		time.Sleep(5 * time.Second)
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()

	go func() {
		defer wg.Done()
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		time.Sleep(5 * time.Second)
		psCertificate.Lock()
		psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("Done")
}
