package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"A", "B", "C", "D"}
	for _, ninja := range evilNinjas {
		// Single go program to attact all 4 ninjas
		// attack(ninja)

		// Multiple processes to attact multiple ninjas at same time - This is called a go routine
		// Goroutine - a task is performed out of the main program execution (asynchronously)
		go attack(ninja)
	}

	// Have to do this, if the main function exits, program execution stops.
	// This can be avoided by using channels
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Trowing stars at target :", target)
	time.Sleep(time.Second * 2)
}
