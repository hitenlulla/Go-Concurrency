// When routines are interdependent, Go has inbuilt feature to send signals between routines that are waiting for each other.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// getReadyForMission()		 // -> We are now ready, After 11372871792 work intervals
	getReadyForMissionWithCond() // -> We are now ready, After 1 work intervals
	broadcastStartOfMission()
}

var ready bool

func getReadyForMission() {
	go getReady()
	workIntervals := 0
	for !ready {
		workIntervals++
	}
	fmt.Printf("We are now ready, After %d work intervals\n", workIntervals)
}

func getReady() {
	sleep()
	ready = true
}

func getReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go getReadyWithCond(cond)
	workIntervals := 0

	// Lock on main process as it is modifying shared workIntervals
	cond.L.Lock()
	for !ready {
		// Wait for the signal from goroutine
		cond.Wait()
		workIntervals++
	}

	// Unlock after modifying
	cond.L.Unlock()
	fmt.Printf("We are now ready, After %d work intervals\n", workIntervals)
}

func getReadyWithCond(cond *sync.Cond) {
	sleep()
	ready = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}

func broadcastStartOfMission() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)
	standByForMission(
		func() {
			fmt.Println("Ninja 1 starting Mission")
			wg.Done()
		},
		cond)
	standByForMission(
		func() {
			fmt.Println("Ninja 2 starting Mission")
			wg.Done()
		},
		cond)
	standByForMission(
		func() {
			fmt.Println("Ninja 3 starting Mission")
			wg.Done()
		},
		cond)
	cond.Broadcast()
	wg.Wait()
	fmt.Println("All ninjas are starting their mission")
}

func standByForMission(fn func(), cond *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fn()
	}()
	wg.Wait()
}
