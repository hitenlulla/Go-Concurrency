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

	evilNinja := "Tom"

	// To avoid abrupt conclusion of main process, we can use channels.
	// Channel is a medium of communication between a routine and a main process.
	// A routine can send msg through a channel which should be intercepted by main process
	// After getting the msg, main process will terminate
	smokeSignal := make(chan bool)
	go attack(evilNinja, smokeSignal)
	// Intercepting msg from channel
	<-smokeSignal

	// Note the concept of channel is only possible between one routine and main.
	// To handle communication between multiple routines and main - we can use bufferred channels
}

func attack(target string, isAttacked chan bool) {
	fmt.Println("Trowing stars at target :", target)
	time.Sleep(time.Second * 1)
	// Sending msg to the channel
	isAttacked <- true
}
