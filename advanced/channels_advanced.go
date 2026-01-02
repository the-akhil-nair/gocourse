package advanced

import (
	"fmt"
)

func main() {

	//variable := make(chan type) '<-' operator
	chan_int := make(chan int)

	// Case 1: sender need receiver to be ready else it blocks the sender
	// As soon as sender sends the data, receiver receives fail to immediately receive it.
	// This will results in deadlock as there is no receiver to receive the data being sent.

	// chan_int <- 1
	// receiver := <-chan_int
	// fmt.Println("Received:", receiver)

	// Case 2: receiver need sender to be ready else it blocks the receiver
	// As soon as receiver is ready to receive the data, sender needs to sends it.
	// This will results in deadlock as there is no sender to send the data.

	// reciever := <-chan_int
	// fmt.Println("Received:", reciever)

	// Case 3: Using goroutine to unblock the sender/receiver
	// As soon as sender sends the data, receiver receives it and both are unblocked.
	go func() {
		chan_int <- 1
	}()

	receiver := <-chan_int
	fmt.Println("Received:", receiver)

	// Case 4: Using goroutine to unblock the sender/receiver
	// As soon as receiver is ready to receive the data, sender sends it and both are unblocked.
	// Receiver sits inside goroutine so sender channel wait till the goroutine is executed.
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Received:", <-chan_int)
	// }()

	// chan_int <- 1

	// Case 5: Using goroutine to unblock both sender and receiver
	// In this case both sender and receiver are inside goroutines.
	// As soon as sender sends the data, receiver receives it and both are unblocked.
	// Channel will wait goroutines to execute before main goroutine exits.
	// go func() {
	// 	fmt.Println("Received:", <-chan_int)
	// }()

	// go func() {
	// 	chan_int <- 1
	// 	fmt.Println("Sent 1 to channel")
	// }()

	// Without this main goroutine will exit before other goroutines complete their execution
	// time.Sleep(3 * time.Second)

	// Case 6: Goroutine to unblock receiver only
	// In this case only receiver will wait for the goroutine to send data.
	// As soon as receiver is ready to receive the data, and come to sender dint send the data, receiver blocks and deadlock occurs.

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// }()

	// receiver := <-chan_int
	// fmt.Println("Received:", receiver)

	// Case 7: Goroutine to unblock sender only
	// In this case only sender will wait for the goroutine to send data.
	// As soon as sender sends data expects receiver to receive it and since goroutine is delaying the receiver, sender wait for all goroutine to end and deadlock occurs.

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// }()

	// chan_int <- 1

	fmt.Println("End of program.")

}
