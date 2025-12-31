package advanced

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			// If the lock is not present and 4 goroutines are incrementing at same time,
			// the final counter value will be less than expected.
			// Suppose initially counter = 0
			// Goroutine 1 reads counter = 0
			// Goroutine 2 reads counter = 0
			// Goroutine 1 increments to 1 and writes back
			// Goroutine 2 increments its read value (0) to 1 and writes back
			// Final counter should be 2 but is 1 due to race condition
			mu.Lock()
			// mu.Lock() and mu.Unlock() marks critical section.
			// Keep the logic minimal inside critical section to avoid performance bottlenecks.
			// Contention occurs when multiple goroutines try to acquire the lock simultaneously,
			// leading to increased waiting times and reduced concurrency.
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counter := &counter{}
// 	numGoroutines := 10

// 	// wg.Add(numGoroutines)
// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 				// counter.count++
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final counter value: %d\n", counter.getValue())

// }
