package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

// When we do ac.count we are directly accessing the count field not through a memory address.
// So we need &ac.count to get the address of the count field for atomic operations.
func (ac *AtomicCounter) increment() {
	// Atomically increments the counter by 1 safely across multiple goroutines.
	// It is faster than using mutexes for simple operations like incrementing a counter.
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	// value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// value++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())
	// fmt.Printf("Final counter value: %d\n", value)
}
