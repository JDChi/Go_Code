package once

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func TestOnce() {
	var once sync.Once

	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1) // will print

	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2) // will not print
}

var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	if conn != nil {
		return conn
	}

	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

// ImplementOnceByOnlyMutex
// if we only use mutex, that means when we call getConn() every time, it always competes the lock to visit conn
func ImplementOnceByOnlyMutex() {
	conn := getConn()
	if conn == nil {
		println("conn is nil")
	}

}

type AtomicOnce struct {
	done uint32
}

// Do
// given two simultaneous calls, the winner of the cas would
// call f, and the second would return immediately, without
// waiting for the first's call to f to complete.
// This is why the slow path falls back to a mutex, and why
// the atomic.StoreUint32 must be delayed until after f returns.
func (o *AtomicOnce) Do(f func()) {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}
