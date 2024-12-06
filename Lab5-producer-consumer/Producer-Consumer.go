// Author: Luka Brennan
// Created on 15/11/2024
// Final edit on 16/11/2024
// Modified: Luka Brennan
// Issue:
// The code template belongs to Joseph Kehoe, The template did not have proper Producer Consumer implementation . This has been fixed in this version

package main

import (
	"fmt"
	"sync"
	"time"
)

// The Producer is a function that will generate 5 integers (0-4) and sends them to the channel
func producer(thelock *sync.Mutex, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		thelock.Lock() // locking shared resources
		channel <- i   // Value for i is sent to the channel
		time.Sleep(time.Millisecond)
		thelock.Unlock() // The lock is released
	}
}

// The Consumer "consumes" the integers that are in the channels until it is closed
func consumer(channel chan int, complete chan bool) {
	for value := range channel { // Retrives values (i) from the channel
		fmt.Println("consumer has:", value)
	}
	complete <- true // Notify main that consumer is done
}

func main() {
	var thelock sync.Mutex
	channel := make(chan int, 5)
	done := make(chan bool)
	var wg sync.WaitGroup

	threads := 5    // The number of producers
	wg.Add(threads) // Add count for producers

	// Start producer
	for i := 0; i < threads; i++ {
		go producer(&thelock, channel, &wg)
	}
	// Start consumer
	go consumer(channel, done)
	// Waiting for all producers to be complete
	wg.Wait()
	close(channel)
	<-done
}
