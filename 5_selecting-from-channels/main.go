package main

import (
	"fmt"
)

func main() {
	ninja1 := make(chan string)
	ninja2 := make(chan string)

	go electNinja(ninja1, "Ninja 1")
	go electNinja(ninja2, "Ninja 2")

	// Here we are selecting a channel, from where the message is coming
	// Note, select statement will be till one channel sends a message
	// Goal is to select only one channel (Channel which is sending the message first)
	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

	roughlyFair()
}

func electNinja(channel chan string, message string) {
	// time.Sleep(2 * time.Second) // To see blocking nature of select
	channel <- message
}

func roughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var n1count, n2count int
	for i := 1; i <= 1000; i++ {
		select {
		case <-ninja1:
			n1count++
		case <-ninja2:
			n2count++
		}
	}

	fmt.Println("Ninja1Count", n1count, "Ninja2Count", n2count)
}
