package advanced

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d at %v\n", req.numTickets, req.personID, req.cost, time.Now())
		// simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}

}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)
	numWorkers := 3

	// start ticket processor/worker
	for range numWorkers {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}
}

// ============= BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	// Whenever a task is received, process it by available worker
// 	// Worker 2 processing task 0 2025-12-30 17:14:18.4100662 +0530 IST m=+0.002332001
// 	// Worker 0 processing task 1 2025-12-30 17:14:18.4100662 +0530 IST m=+0.002332001
// 	// Worker 1 processing task 2 2025-12-30 17:14:18.4100662 +0530 IST m=+0.002332001
// 	// As soon as any worker is free, it picks up next task, In below case 0.
// 	// Worker 0 processing task 3 2025-12-30 17:14:19.4157183 +0530 IST m=+1.007972601
// 	// Result: 2
// 	// Result: 0
// 	// Result: 4
// 	// Worker 2 processing task 4 2025-12-30 17:14:19.4157183 +0530 IST m=+1.007972601
// 	// Since tasks is recieve only channel, next recieve will picked by any available worker
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d %v\n", id, task, time.Now())
// 		// Simulate work
// 		time.Sleep(time.Second)
// 		// Since worker is doing some processing, it sends result to results channel after processing
// 		results <- task * 2
// 	}
// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// Create workers
// 	for i := range numWorkers {
// 		// At any point of time, only numWorkers goroutines will be processing tasks concurrently
// 		go worker(i, tasks, results)
// 	}

// 	// Send values to the tasks channel
// 	for i := range numJobs {
// 		// This loop will send all tasks to the tasks channel quickly
// 		// and then workers will pick them up one by one as they become free
// 		tasks <- i
// 	}

// 	close(tasks)

// 	// Collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// }
