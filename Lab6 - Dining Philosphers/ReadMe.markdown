# Dining Philosophers Problem

This project demonstrates the **Dining Philosophers Problem** implemented in Go, developed by **Joseph Kehoe**. The solution uses **channels**, **goroutines**, and **semaphores** to manage resource contention among philosophers. However, the current implementation has **critical issues**, including cases where **three or more philosophers eat simultaneously**, violating the problem's constraints.

## Known Issues

**Concurrent Eating**:
   - The current semaphore logic, meant to restrict the number of philosophers eating at a time, fails occasionally.
   - This leads to situations where three or more philosophers eat simultaneously, breaking the rules of the Dining Philosophers Problem.

## How to Run the Code

1. Download the zip folder and save it someware - then unzip it
2. Open a terminal and navigate to the file's directory where the saved file is.
3. Run the file using:
   go run "Lab6-DiningPhilosophers/dinPhili(1).go-"
   or cd (on windows) to the Lab6-DiningPhilosophers/dinPhili(1).go

## Licensing
<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><span property="dct:title">Concurrent Development - Labs</span> by <span property="cc:attributionName">Joseph Kehoe</span> is licensed under <a href="https://www.gnu.org/licenses/gpl-3.0.html" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">GPL-3.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt=""></a></p>
