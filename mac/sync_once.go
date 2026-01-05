package mac

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	println("Initialization function executed")
}

func main() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
			fmt.Println("Goroutine started")

		}()
	}
	wg.Wait()
}
