package advanced

import (
	"fmt"
	"time"
)

// func main() {
// 	// =========== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of program.")
// }

func main() {
	// buffered channels allow senders to continue working without blocking until the buffer is full.
	// We can manage the rate of data transfer between producer and consumer and act as a load balancer.
	// Unbuffered channel need immediate receiver to be ready to receive the data, else it blocks the sender.
	// ================== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	// make(chan Type, capacity)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		// If this line is commented out, the main goroutine will block forever on the next send operation and hence deadlock occurs
		// In the below line as soon as receive happens, the send in main goroutine unblocks and there is a chance that println will not print.
		fmt.Println("Received:", <-ch) //ends <- starts
	}()
	fmt.Println("Blocking starts")
	ch <- 3 // Blocks because the buffer is full and waiting go routine to receive
	//fmt.Println("Blocking ends")
	//fmt.Println("Received:", <-ch)
	//fmt.Println("Received:", <-ch)
}
