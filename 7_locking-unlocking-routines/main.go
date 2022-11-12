package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Suppose one routine is dependent on result from another routine. That routine needs to be blocked till the dependency is completed.
To block and unblock the routine, we can use MUTEX.
A MUTEX is a token that allows the routine to run. Routine that holds the mutex will have the rights to execute -> (Locking).
The moderator (Who changes the mutex) is the Operating system.
When the mutex goes back to the moderator from the routine, it means the task was completed -> (Unlocking)

This is used to make sure that on any given time, only one routine is making changes to the resource.
*/

var count int
var lock sync.Mutex
var rwLock sync.RWMutex

func increment() {
	// count++ is not atomic
	// It is roughly translated to
	// temp := count
	// temp = temp + 1
	// count = temp

	// Hence when the routine is called, it might happen that some of the increments get skipped hence we never get 1000
	// This can be fixed by locking this function and then unlocking it
	lock.Lock()
	count++
	lock.Unlock()
}

func mutexDemo() {
	itr := 1000
	for i := 0; i < itr; i++ {
		// Here we can see that count never becomes 1000
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Result Count is", count)

	/*
		In the above example: count is a shared resource between all increment() go routines.
		Without MUTEX:
			While modifying the count, Some of the increment() routines wer not completed hence we never got 1000.
		with MUTEX:
			We make sure that every go routine is individually completed by locking and then unlocking at end of function.
			Hence we got 1000.
			When the lock is made. No other go routine will modify the count until it is unlocked
	*/
}

func read() {
	rwLock.RLock()
	defer rwLock.RUnlock()

	fmt.Println("Read Locking")
	// time.Sleep(1 * time.Second)
	fmt.Println("Counter:", count)
	fmt.Println("Read unlocking")
}

func write() {
	rwLock.Lock()
	defer rwLock.Unlock()

	fmt.Println("Write Locking")
	// time.Sleep(1 * time.Second)
	count++
	fmt.Println("Counter Incremented to:", count)
	fmt.Println("Write unlocking")
}

/*
A ReadWrite MUTEX is a mutex where the lock can be aquired by multiple read routines but only one write routine at a time.
Suppose Routines need to read the shared information, multiple routines can do that at a time.
But if a routine needs to edit the shared information, only one routine can do that at a time.
*/

func rwMutexDemo() {
	go read()
	go write()
	go read()

	time.Sleep(7 * time.Second)
	fmt.Println("Done")
}

func main() {
	// mutexDemo()
	rwMutexDemo()
}
