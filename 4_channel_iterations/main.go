package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	// Throwing 1 star
	// go throwNinjaStar(channel)
	// fmt.Println(<-channel)

	// Throwing 3 stars - one by one
	// count := 3
	// go throwNinjaStar(channel, count)
	// Iterating over channel - for loop
	// for i := 1; i <= count; i++ {
	// 	fmt.Println(<-channel)
	// }

	/*
		// Iterating over channel - for range
		// Count is considered in throwNinjaStar
		// If using range - not that if all the msgs of channel are printed, range tries to find the next message in channel and throws deadlock
		// To fix this make sure to close the channel once all the messages are sent
	*/
	// go throwNinjaStar(channel)
	// for message := range channel {
	// 	fmt.Println(message)
	// }

	// Specific for loop for channel iteration - consists of message and open flag. (Make sure to close the channel)
	go throwNinjaStar(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)

	}
}

// for-loop
// func throwNinjaStar(channel chan string, count int) {
// 	rand.Seed(time.Now().UnixNano())
// 	for j := 1; j <= count; j++ {
// 		score := rand.Intn(10)
// 		channel <- fmt.Sprint("You scored: ", score)
// 	}
// }

// for range & for message,open
func throwNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	count := 3
	for j := 1; j <= count; j++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
	// Closing the channel
	close(channel)
}
