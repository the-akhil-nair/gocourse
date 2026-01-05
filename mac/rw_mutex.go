package mac

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func ReadCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Multiple goroutines can hold the read lock simultaneously
	// Read locks block write locks but not other read locks
	rwmu.RLock()
	fmt.Printf("Goroutine %d: Read Counter Value: %d at %s\n", id, counter, time.Now())
	rwmu.RUnlock()

}

func WriteCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	// One goroutine can hold the write lock at a time, blocking all other reads and writes
	// While a write lock is held, no other goroutine can acquire a read or write lock
	rwmu.Lock()
	// The below line shows that write lock blocks reads
	time.Sleep(2 * time.Second) // Simulate write delay
	counter = value
	fmt.Printf("Wrote Counter Value: %d at %s\n", counter, time.Now())
	rwmu.Unlock()
}

func main() {
	// Sample Output with prolonged write lock:
	// First goRoutine triggered.
	// Goroutine 0: Read Counter Value: 0 at 2026-01-02 15:25:23.9631841 +0530 IST m=+0.001064001
	// WriteCounter triggered and acquired write lock for 2 seconds.
	// Wrote Counter Value: 20 at 2026-01-02 15:25:25.9715949 +0530 IST m=+2.009459701
	// All other read goroutines blocked until write is done.
	// Goroutine 4: Read Counter Value: 20 at 2026-01-02 15:25:25.9727645 +0530 IST m=+2.010629301
	// Goroutine 2: Read Counter Value: 20 at 2026-01-02 15:25:25.9728756 +0530 IST m=+2.010740401
	// Goroutine 3: Read Counter Value: 20 at 2026-01-02 15:25:25.9729299 +0530 IST m=+2.010794701
	// Goroutine 1: Read Counter Value: 20 at 2026-01-02 15:25:25.9735432 +0530 IST m=+2.011408001
	// Final Counter Value: 20

	//  Sample Output without prolonged write lock:
	// All ReadCounter goroutines complete before WriteCounter starts signifies that read locks do not block each other.
	// Goroutine 3: Read Counter Value: 0 at 2026-01-02 15:25:00.4320469 +0530 IST m=+0.001055501
	// Goroutine 4: Read Counter Value: 0 at 2026-01-02 15:25:00.4320469 +0530 IST m=+0.001055501
	// Goroutine 1: Read Counter Value: 0 at 2026-01-02 15:25:00.4320469 +0530 IST m=+0.001055501
	// Goroutine 2: Read Counter Value: 0 at 2026-01-02 15:25:00.4320469 +0530 IST m=+0.001055501
	// Goroutine 0: Read Counter Value: 0 at 2026-01-02 15:25:00.4320469 +0530 IST m=+0.001055501
	// Wrote Counter Value: 20 at 2026-01-02 15:25:03.4525242 +0530 IST m=+3.021510101
	// Final Counter Value: 20
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go ReadCounter(i, &wg)
	}

	wg.Add(1)
	// time.Sleep(1 * time.Second) // Ensure reads happen before write
	go WriteCounter(&wg, 20)

	wg.Wait()
	fmt.Println("Final Counter Value:", counter)
}
