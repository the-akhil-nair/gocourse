package mac

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	go func() {
		mu1.Lock()

		fmt.Println("Goroutine 1 acquired mu1", time.Now())
		time.Sleep(1 * time.Second)
		// Trying to acquire mu2 which is held by goroutine 2
		// Because both goroutines are waiting for each other to release the locks, a deadlock occurs here
		// Because mu2 is held by goroutine 2, both goroutines are stuck waiting indefinitely and next line will never be reached
		mu2.Lock()
		fmt.Println("Goroutine 1 acquired mu2", time.Now())
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 acquired mu1", time.Now())
		time.Sleep(1 * time.Second)
		// Trying to acquire mu1 which is held by goroutine 1
		mu2.Lock()
		// To avoid deadlock, uncomment the next line
		// mu2.Lock()
		fmt.Println("Goroutine 2 acquired mu2", time.Now())
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
	//select {}
}
