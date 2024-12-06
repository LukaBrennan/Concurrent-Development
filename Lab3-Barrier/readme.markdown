# Simple Barrier

This project demonstrates a **Simple Barrier** mechanism implemented in Go, developed by **Luka Brennan**, as a solution to synchronize multiple goroutines. The implementation utilizes **mutexes** and **semaphores** to coordinate concurrent tasks. This barrier ensures all goroutines reach a designated synchronization point (`Part A`) before proceeding to the next (`Part B`).

## Implementation Details

The `Barrier.go` code features the following key components:

- **Goroutines and WaitGroup**:
  - Each goroutine represents a concurrent task, coordinated using a `sync.WaitGroup` to ensure all tasks complete before the program exits.

- **Mutex for Critical Sections**:
  - A `sync.Mutex` protects the shared variable `arrived` from concurrent access, ensuring thread safety when updating or checking the count of arrived goroutines.

- **Semaphore for Resource Management**:
  - A semaphore (`semaphore.Weighted`) ensures the barrier logic synchronizes access among all goroutines.

### Synchronization Process

1. **Part A Execution**:
   - Each goroutine executes its task in `Part A`. Once it reaches the synchronization point, it updates the `arrived` count in a critical section controlled by a mutex.

2. **Barrier Check**:
   - The goroutines wait until all `arrived` values are synchronized, mimicking a barrier. A semaphore helps manage the synchronization mechanism effectively.

3. **Part B Execution**:
   - After synchronization, the goroutines proceed to `Part B`, completing the remaining tasks.

   ## How to Run the Code

1. Download the zip folder and save it someware - then unzip it
2. Open a terminal and navigate to the file's directory where the saved file is.
3. Run the file using:
   go run "Lab3-Barrier/simplebarrier.go"


## Licensing

<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><span property="dct:title">Concurrent Development - Simple Barrier</span> by <span property="cc:attributionName">Luka Brennan</span> is licensed under <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY-NC-SA 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1" alt=""></a></p>
