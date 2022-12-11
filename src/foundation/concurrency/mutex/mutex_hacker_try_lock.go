package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识 waiter 的起始 bit 位置
)

// TryLockMutex 一个带有 TryLock 的 Mutex
// From Go 1.18 Mutex has a new function TryLock
type TryLockMutex struct {
	sync.Mutex
}

func (m *TryLockMutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func (m *TryLockMutex) Count() int {
	// get the value of the state
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	// 得到等待者的数值
	v = v >> mutexWaiterShift
	// 加上锁持有者的数量，0 或者 1
	v = v + (v & mutexLocked)
	return int(v)
}

func (m *TryLockMutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

func (m *TryLockMutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

func (m *TryLockMutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

func TryLockMutexTest() {
	var tryLockMutex TryLockMutex
	go func() {
		tryLockMutex.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		tryLockMutex.Unlock()
	}()

	time.Sleep(time.Second)

	ok := tryLockMutex.TryLock()
	if ok {
		fmt.Println("got the lock")
		// do something
		tryLockMutex.Unlock()
		return
	}
	fmt.Println("can't get the lock")
}

func TryLockMutexTest2() {
	var tryLockMutex TryLockMutex
	for i := 0; i < 1000; i++ {
		go func() {
			tryLockMutex.Lock()
			time.Sleep(time.Second)
			tryLockMutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("waiting: %d. isLocked: %t, woken: %t, starving: %t\n", tryLockMutex.Count(), tryLockMutex.IsLocked(), tryLockMutex.IsWoken(), tryLockMutex.IsStarving())
}
