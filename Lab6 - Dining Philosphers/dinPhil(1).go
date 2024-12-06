// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
// GPL Licence
// MISSING:
// 1. Readme
// 2. Full licence info.
// 3. Comments
// 4. It can Deadlock!
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate thinking by a philosopher
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.Intn(5))
	time.Sleep(X * time.Second) // Wait random time
	fmt.Println("Phil: ", index, "was thinking")
}

// Simulate eating by a philosopher
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.Intn(5))
	time.Sleep(X * time.Second) // Wait random time
	fmt.Println("Phil: ", index, "was eating")
}

// Philosophers' behavior
func doPhilStuff(index int, wg *sync.WaitGroup, leftFork, rightFork chan struct{}, eatSemaphore chan struct{}) {
	defer wg.Done()
	for i := 0; i < 2; i++ { // Each philosopher eats twice
		think(index)

		// Limit the number of philosophers trying to eat concurrently
		eatSemaphore <- struct{}{} // Acquire eating permission
		<-leftFork                 // Acquire the left fork
		<-rightFork                // Acquire the right fork

		eat(index)

		leftFork <- struct{}{}  // Release the left fork
		rightFork <- struct{}{} // Release the right fork
		<-eatSemaphore          // Release eating permission
	}
}

func main() {
	philCount := 5
	var wg sync.WaitGroup
	wg.Add(philCount)

	// Initialize forks as channels
	forks := make([]chan struct{}, philCount)
	for i := 0; i < philCount; i++ {
		forks[i] = make(chan struct{}, 1) // Buffered channel with size 1
		forks[i] <- struct{}{}            // Each fork is initially available
	}

	// Semaphore to limit the number of philosophers eating concurrently
	eatSemaphore := make(chan struct{}, 2) // Allow at most 2 philosophers to eat at a time

	// Create philosophers
	for i := 0; i < philCount; i++ {
		leftFork := forks[i]
		rightFork := forks[(i+1)%philCount]
		go doPhilStuff(i, &wg, leftFork, rightFork, eatSemaphore)
	}

	wg.Wait() // Wait for all philosophers to finish
}
