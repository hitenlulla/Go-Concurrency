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

	// channel := make(chan string)
	// channel <- "Hello Channel"
	/* The above code creates a deadlock
	Why?- Channels are not variables. They don't have a capacity [by default] to store the msgs sent to it.
	If a msg is sent from a go routine to a channel, it should be intercepted by some other go routine immediately.
	To solve this, we can give capacity to the channel - this is called a buffered channel.
	It is used to store the msg in buffer and send it FIFO (Queue).
	*/

	channel := make(chan string, 2)
	channel <- "First Value"
	channel <- "Second Value"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
