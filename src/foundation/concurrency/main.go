package main

import "Go_Code/src/foundation/concurrency/once"

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
	//rwmutex.TestRWMutex()

	// waitgroup
	//waitgroup.CommonUse()
	// waitgroupmistake
	//waitgroup.MistakeAdd()
	//waitgroup.MistakeDone()
	//waitgroup.MistakeAddTiming()
	//waitgroup.MistakeReuseWaitGroupBeforeDone()

	// once
	//once.TestOnce()
	//once.MistakeDeathLock()
	once.MistakeUnInit()
}
