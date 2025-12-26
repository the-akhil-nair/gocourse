package advanced

import (
	"fmt"
	"time"
)

func channels() {

	//variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	// Its trying to send value to channel but there is no goroutine ready. A goroutine is needed to receive the value.
	// Channels in go are blocking by default. Hence the below line will block the main goroutine forever creating a deadlock and need to be added under a separate goroutine.
	// greeting <- greetString
	// receiver := <-greeting
	// fmt.Println(receiver)

	go func() {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	// Below channel receives since it is part of main goroutine which is always running and communication happens between two goroutines.
	// Below goroutine will block until there is a value to receive from channel else end up in deadlock.
	// Once receiver receives value from channel, the sending goroutine is unblocked and able to send next value.
	// These issues are run time issues and may not be visible during compile time.
	receiver := <-greeting
	fmt.Println(receiver)
	// Need to call as many times as values are sent to channel.
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")

}
