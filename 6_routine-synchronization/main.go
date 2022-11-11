package main

import (
	"fmt"
	"sync"
)

func main() {
	evilNinjas := []string{"Tom", "Dom", "Ron"}
	// Initialise a wait counter
	var beeper sync.WaitGroup
	// Add number of counts it has to wait
	beeper.Add(len(evilNinjas))
	// Call routine and pass reference to beeper (reference is needed because the counter will be decremented at every beeper.Done())
	for _, ninja := range evilNinjas {
		go attack(ninja, &beeper)
	}

	// Wait for all routines to finish
	beeper.Wait()
	// End of main()
	fmt.Println("Attack succesfull")
}

func attack(ninja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked Ninja:", ninja)
	// Declare that this routine was completed
	beeper.Done()
}
