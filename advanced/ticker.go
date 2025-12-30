package advanced

import (
	"fmt"
	"time"
)

func populateCache(t time.Time) {
	fmt.Println("Populating cache at:", t)
}

func fetchUpdates(t time.Time) {
	fmt.Println("Fetching updates at:", t)
}

func cleanupResources(t time.Time) {
	fmt.Println("Cleaning up resources at:", t)
}

func main() {

	ticker1 := time.NewTicker(2 * time.Second)
	defer ticker1.Stop()

	ticker2 := time.NewTicker(3 * time.Second)
	defer ticker2.Stop()

	ticker3 := time.NewTicker(5 * time.Second)
	defer ticker3.Stop()

	timer := time.After(10 * time.Second)

	for {
		select {
		case t1 := <-ticker1.C:
			populateCache(t1)
		case t2 := <-ticker2.C:
			cleanupResources(t2)
		case t3 := <-ticker3.C:
			fetchUpdates(t3)
		case <-timer:
			fmt.Println("Stopping all periodic tasks.")
			return
		}
	}
}

// func main() {

// 	ticker := time.NewTicker(time.Second)
// 	stop := time.After(5 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			fmt.Println("Tick at:", tick)
// 		case <-stop:
// 			fmt.Println("Stopping ticker.")
// 			return
// 		}
// 	}
// }

// ========= SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }
// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	// for tick := range ticker.C {
// 	// 	i *= 2
// 	// 	fmt.Println(tick)
// 	// }
// }
