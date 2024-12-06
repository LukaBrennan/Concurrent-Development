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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Luka Brennan
// Issues:
// The barrier is not implemented!
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived int, wg *sync.WaitGroup, thelock *sync.Mutex, sem *semaphore.Weighted, ctx context.Context) bool { // add *arrived and semaphoere.waited , max int

	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	//for i := 0; i < goNum; i++ {
	if arrived == 0 {
		thelock.Lock()
		time.Sleep(time.Nanosecond)
		arrived++
		sem.Acquire(ctx, 1)
		sem.Release(1)
		thelock.Unlock()

	} else {
		thelock.Lock()
		time.Sleep(time.Nanosecond)
		arrived++
		sem.Acquire(ctx, 1)
		sem.Release(1)
		thelock.Unlock()
	}
	//}
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	arrived := 0
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, arrived, &wg, &theLock, sem, ctx) // add &arrived, &sem
	}
	wg.Wait() //wait for everyone to finish before exiting
}
