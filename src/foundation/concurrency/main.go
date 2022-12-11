package main

import "Go_Code/src/foundation/concurrency/mutex"

func main() {
	// mutex
	//WithoutMutex()
	//FixWithMutex()

	// mutex mistake
	//MistakeNoLockUnlockPairing()
	//MistakeCopyUsedMutex()
	//mutex.MistakeReentrantLock()
	mutex.MistakeDeadLock()
}
