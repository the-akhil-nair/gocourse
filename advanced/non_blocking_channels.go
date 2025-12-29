package advanced

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)

	// === NON BLOCKING RECEIVE OPERATION
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No messages available.")
	// }

	// // === NON BLOCKING SEND OPERATION
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message.")
	// default:
	// 	fmt.Println("Channel is not ready to receive.")
	// }

	// === NON BLOCKING OPERATION IN REAL TIME SYSTEMS

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				// Stopping may not come as main goroutine exits before this goroutine gets chance to run.
				// Adding sleep in default case may help in getting stopping message printed.
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				// Adding 500 millisecond sleep not printing stopping message as well as printing waiting for data message 3 times.
				// Again its related to how faster happening concurrently behind the scenes.
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
	fmt.Println("Program ended.")

}
