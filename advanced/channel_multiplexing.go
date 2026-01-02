package advanced

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		// close(ch)
		// ch <- 2 // panic: send on closed channel
	}()

	go func() {
		// Simulate some work
		// time.Sleep(1 * time.Second)
		ch <- 2
		// close(ch)
	}()

	// time.Sleep(1 * time.Second)
	// close(ch)

	for {
		time.Sleep(500 * time.Millisecond)
		select {
		// Channel needs to be closed and empty then only ok will be false
		case msg, ok := <-ch:
			// If this condition is not added, then the infinite loop will continue trying to read from closed channel
			if !ok {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		default:
			fmt.Println("No channels ready...Closing channel")
			close(ch)
			return
		}
	}
}

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received:", msg)
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout.")
// 	}
// }

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {

// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()
// 	go func() {

// 		// time.Sleep(time.Second)
// 		ch2 <- 2
// 	}()

// 	time.Sleep(2 * time.Second)

// 	for range 2 {
// This is done to check the conditions of both channels multiple times
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Received from ch1:", msg)
// 		case msg := <-ch2:
// 			fmt.Println("Received from ch2:", msg)
// 		default:
// 			fmt.Println("No channels ready...")
// 		}
// 	}

// 	fmt.Println("End of program")

// }
