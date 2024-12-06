//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Luka Brennan
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
// 1. Change mutex to atomic variable
// 2. Make it a reusable barrier
//
//	Edited by Luka Brennan
//	Added:
//	1. Atomic varaibles
//	2. reusable barrier
//
// --------------------------------------------
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// DoStuff function uses a barrier to synchronize multiple goroutines ensuring that all goroutines finish Part A before any proceeds to Part B.
func doStuff(goNum int, arrived *int32, max int, wg *sync.WaitGroup, barrierChan chan bool, iteration int) bool {
	defer wg.Done()

	for i := 0; i < iteration; i++ { // Loop
		fmt.Println("Part A", goNum)
		time.Sleep(time.Second)

		if atomic.AddInt32(arrived, 1) == int32(max) { // Checking if all goroutines have arrived
			for i := 0; i < max-1; i++ { // signals the other goroutines using barrierchan
				barrierChan <- true
			}

			atomic.StoreInt32(arrived, 0) // resetting arrived for the next iteration
		} else {
			<-barrierChan // Wait until the last goroutine signals
		}

		fmt.Println("Part B", goNum)
		time.Sleep(time.Second)
	}
	return true
}

func main() {
	totalRoutines := 3
	iteration := 3
	var arrived int32
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	barrierChan := make(chan bool)
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &arrived, totalRoutines, &wg, barrierChan, iteration)
	}
	wg.Wait()
}
