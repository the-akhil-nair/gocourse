package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan int)

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- 0
// 	}()

// 	<-done
// 	fmt.Println("Finished.")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value is received
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch // Blocking until a value is sent
// 	fmt.Println(value)
// }

// // ========= SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, numGoroutines)

// 	for i := range numGoroutines {
// 		// This line works as there are more sends than receives and hence no deadlock occurs.
// 		// But this will create goroutine leak as the last goroutine will be blocked forever trying to send to done channel.
// 		// for i := range numGoroutines + 1 {

// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // SENDING SIGNAL OF COMPLETION
// 		}(i)
// 	}

// 	// Since done is buffered channel with capacity equal to number of goroutines, we can safely receive from it numGoroutines times without blocking
// 	// So commenting out the below line will not cause deadlock.
// 	// The below line will cause deadlock becauase now receive is running more times than the sends.
// 	// It blocked till all goroutines are done, and at last failed because of more receives than sends.
// 	// for range numGoroutines + 1 {
// 	for range numGoroutines {

// 		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
// 	}

// 	fmt.Println("All goroutines are completed")
// }

// ========== 	SYNCHRONIZING DATA EXCHANGE
func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel closed before Goroutine could send a value to the channel and main program tries to receive from closed channel.

	// Receiving values until channel is closed and since channel is not closed, we will get deadlock.
	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}
