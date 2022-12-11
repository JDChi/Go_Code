package main

import "Go_Code/src/foundation/concurrency/rwmutex"

func main() {
	// mutex
	//WithoutMutex()
	//FixWithMutex()

	// mutex mistake
	//MistakeNoLockUnlockPairing()
	//MistakeCopyUsedMutex()
	//mutex.MistakeReentrantLock()
	//mutex.MistakeDeadLock()
	//mutex.TryLockMutexTest2()
	rwmutex.TestRWMutex()
}
