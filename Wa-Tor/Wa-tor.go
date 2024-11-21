package main

import (
	"fmt"
	"math/rand"
	// The Chronos
)

func main() {
	row := 100
	col := 100
	blue := "\033[34m"
	reset := "\033[0m"
	fish := ">('>"
	water := "_"
	//fishbreedtime := 3

	fishtank := make([][]string, row)
	for i := range fishtank {
		fishtank[i] = make([]string, col)
	}

	for i := range fishtank {
		for j := range fishtank[i] {
			if rand.Intn(2) == 0 {
				fishtank[i][j] = blue + fish + reset
			} else {
				fishtank[i][j] = water
			}
		}
	}

	for i := range fishtank {
		for j := range fishtank[i] {
			fmt.Print(fishtank[i][j], " ")
		}
		fmt.Println()
	}
}

//Notes for fish
//Need to decalre a breeding time for fish. this will use chronons. once fihs
//Need to poulate the 2d array with fish (for starters). then have those fish be able to move around (using chronos ?)
//Once fish can move around, then test the breeding time. make sure that they can bread and it is updated through the 2d array.
//Need to make the 2d array be able to move. this could be the having different instances of the 2d array or one continues flow. Chronos
// Need to have the fish move in a pattern.
// the chronos should keep the same 100x100 2d array, on each chronos the fish should move to the next aviable empty water. on every third cchronos when a fish moves to the next empty space, ensure that the prvious space will be emptywater so that a new fish can be producded (fish breeds).
// This may require a mutex lock or some sort od locking. once this is completed the timer of the chronos will be reset back to 0.
