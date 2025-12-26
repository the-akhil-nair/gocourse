package advanced

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value
// Goroutines do not stop the program flow and are non blocking
// Goroutines will still be there in memory even if the main thread is finished executing but will be overwritten by the garbage collector

func main() {
	var err error

	// goruntime scheduler manages all the goroutines in the program
	// it allocates time for each goroutine to run and switches between them

	// m:n scheduling model
	// m = number of goroutines
	// n = number of OS threads
	// The Go runtime maps m goroutines onto n OS threads
	// This allows for efficient use of system resources and better performance

	// Multiplexing is like switching, and go scheduler does this automatically

	// When a goroutine performs a blocking operation (like I/O), the scheduler can switch to another goroutine
	// This ensures that the program remains responsive and can handle multiple tasks concurrently
	// Older programs used to have a fixed number of OS threads, but modern Go versions dynamically adjust the number of threads based on workload
	// Task scheduling and task switching is handled by the Go runtime, not the OS.

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()

	// err = go doWork() // This is not accepted

	// there is a chance that printLetters might execute before printNumbers because of the thread scheduling done by the go runtime
	go printNumbers()
	go printLetters()

	// gorouting provides concurrency.
	// It go scheduler manages the execution of goroutines on multiple cores.

	time.Sleep(2 * time.Second)

	// Since doWork is running in a goroutine, we need to wait for it to finish before checking the error
	// In a real-world scenario, we would use sync.WaitGroup or channels to synchronize goroutines
	// If this might have added before time.Sleep(2 * time.Second), err might still be nil here and printed Work completed successfully
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}

	// Infinite loop inside goroutine will result in goroutine leak.
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}
