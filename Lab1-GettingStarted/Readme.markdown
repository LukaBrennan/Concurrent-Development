# Go Concurrency Examples

This file contains multiple examples demonstrating various concurrency mechanisms in Go. Each section highlights a unique approach to solving concurrency-related problems using tools like **atomic variables**, **mutexes**, **semaphores**, and **channels**.

---

## Example 1: Atomic Counter Increment

### Description
This example uses the `atomic.Int64` type to safely increment a shared counter across multiple goroutines. The **atomic operations** ensure thread safety without explicit locking mechanisms.

### Key Concepts:
- **Concurrency Tools Used**: `atomic.Int64`, `sync.WaitGroup`.
- Goroutines safely increment the counter using `atomic.Add`.

### Expected Behavior:
- A total of 10 goroutines increment the counter by 1000 each. The final value is `10000`.

## Example 2: Mutex-Based Counter Increment

### Description
This example protects a shared counter with a **mutex** (`sync.Mutex`). Each goroutine locks the mutex, increments the counter, and then unlocks the mutex.

### Key Concepts:
- **Concurrency Tools Used**: `sync.Mutex`, `sync.WaitGroup`.
- Mutex prevents race conditions by serializing access to the counter.

### Expected Behavior:
- A total of 10 goroutines increment the counter by 1000 each. The final value is `10000`.

## Example 3: Goroutines with Rendezvous Point

### Description
This example introduces random delays for goroutines, followed by a synchronization point (rendezvous) where all goroutines must wait before proceeding.

### Key Concepts:
- **Concurrency Tools Used**: `sync.WaitGroup`.
- Each goroutine performs **Part A**, syncs at a rendezvous, and proceeds to **Part B**.

### Expected Behavior:
- Logs indicate random execution of **Part A** followed by synchronization at **Part B**.

## Example 4: Semaphore-Based Worker Pool

### Description
This example demonstrates a **worker pool** pattern using a semaphore to limit the number of concurrent goroutines. Each goroutine calculates the number of steps to reach 1 under the **Collatz conjecture**.

### Key Concepts:
- **Concurrency Tools Used**: `semaphore.Weighted`.
- The semaphore restricts active worker count to the number of logical processors (`GOMAXPROCS`).

### Expected Behavior:
- Collatz steps for 64 numbers are computed in parallel, respecting the worker pool limit.

## Example 5: Custom Semaphore Implementation

### Description
This example implements a **custom semaphore** using a buffered channel to limit the number of concurrent tasks.

### Key Concepts:
- **Concurrency Tools Used**: `chan struct{}`, `sync.WaitGroup`.
- Tasks acquire and release semaphore slots using the channel.

### Expected Behavior:
- A maximum of 5 tasks can run concurrently, with 20 tasks queued for execution.

## Example 6: Channel-Based Synchronization

### Description
This example uses a **channel as a barrier** to synchronize two goroutines. The channel ensures that **Part A** of one task completes before **Part B** of the other task begins.

### Key Concepts:
- **Concurrency Tools Used**: `chan bool`, `sync.WaitGroup`.
- Channels coordinate task sequencing between two goroutines.

### Expected Behavior:
- Logs show synchronization between **doStuffOne** and **doStuffTwo**.

## How to Run the Code

1. Download the zip folder and save it someware - then unzip it
2. Open a terminal and navigate to the file's directory where the saved file is.
3. Run the file using:
   go run "Lab1-GettingStarted/then the go code you want.go"

## Licensing
<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><span property="dct:title">Concurrent Development - Labs</span> by <span property="cc:attributionName">Luka Brennan</span> is licensed under <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY-NC-SA 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1" alt=""></a></p>
