package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Suppose there are 100 routines that need to perform some task, but only one of them being sucessful is enough.
Say, the 10th routine was succesful then there is no purpose of running remaining 90 routines, to mark the status as comleted.
This can be done using the sync.Once
*/
var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			if isTreasureFound() {
				// This makes sure that if the mission is completed it stops
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	checkMissionCompleted()
}

func checkMissionCompleted() {
	if missionCompleted {
		fmt.Println("Mission is completed")
	} else {
		fmt.Println("Mission is failure")
	}
}

func markMissionCompleted() {
	// Here we can see the mission was completed, yet all the routines continues to run.
	// To avoid this use sync.Once
	fmt.Println("Marking Mission completed")
	missionCompleted = true
}

func isTreasureFound() bool {
	rand.Seed(time.Now().UnixNano())
	// only 1 out of 10 routines will pass.
	if rand.Intn(10) == 0 {
		return true
	} else {
		return false
	}
}
