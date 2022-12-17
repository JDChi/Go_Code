package once

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
)

func MistakeDeathLock() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
			fmt.Println("init")
		})
	})
}

func MistakeUnInit() {
	var once sync.Once
	var googleConn net.Conn

	once.Do(func() {
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})
	googleConn.Write([]byte("Get / HTTP/1.1\r\nHost: google.com\r\n Accept: */"))
	io.Copy(os.Stdout, googleConn)
}

// OnceWithErr
/// make f() return err to fix un init mistake
type OnceWithErr struct {
	m    sync.Mutex
	done uint32
}

func (o *OnceWithErr) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.doSlow(f)
}

func (o *OnceWithErr) doSlow(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		// change the value only if err == nil
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}
