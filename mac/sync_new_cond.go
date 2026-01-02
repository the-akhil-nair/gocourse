package mac

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type Buffer struct {
	mutex sync.Mutex
	cond  *sync.Cond
	items []int
}

func newBuffer(size int) *Buffer {
	b := &Buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Buffer) Produce(item int) {
	// We need to lock the mutex before accessing the buffer and using the condition variable
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for len(b.items) == bufferSize {
		// When producer is faster than consumer, this condition will be true first.
		// Once the buffer is full, the producer has to wait for the consumer to consume an item
		fmt.Println("Producer is waiting at", time.Now())
		// Release the lock temporarily and wait for a signal
		b.cond.Wait() // Wait until there is space putting this goroutine to sleep and releasing the lock
		// Consumer is waiting at 2026-01-02 17:47:51.1302976 +0530 IST m=+0.001641401
		// Produced: 100 2026-01-02 17:47:51.1308886 +0530 IST m=+0.002232401
		// Producer signaled at 2026-01-02 17:47:51.1308886 +0530 IST m=+0.002232401
		// Consumer woke up at 2026-01-02 17:47:51.1314687 +0530 IST m=+0.002812501
		// Consumed: 100 2026-01-02 17:47:51.1314687 +0530 IST m=+0.002812501
		// Consumer signaled at 2026-01-02 17:47:51.1314687 +0530 IST m=+0.002812501
		// Produced: 101 2026-01-02 17:47:51.1341541 +0530 IST m=+0.005497801
		// Producer signaled at 2026-01-02 17:47:51.1341541 +0530 IST m=+0.005497801
		// Produced: 102 2026-01-02 17:47:51.1496217 +0530 IST m=+0.020965301
		// Producer signaled at 2026-01-02 17:47:51.1496217 +0530 IST m=+0.020965301
		// Produced: 103 2026-01-02 17:47:51.16509 +0530 IST m=+0.036433501
		// Producer signaled at 2026-01-02 17:47:51.1657323 +0530 IST m=+0.037075801
		// Produced: 104 2026-01-02 17:47:51.1808419 +0530 IST m=+0.052185301
		// Producer signaled at 2026-01-02 17:47:51.1808419 +0530 IST m=+0.052185301
		// Produced: 105 2026-01-02 17:47:51.1972866 +0530 IST m=+0.068629801
		// Producer signaled at 2026-01-02 17:47:51.1972866 +0530 IST m=+0.068629801 // Buffer is full now, next produce will wait
		// Producer is waiting at 2026-01-02 17:47:51.2129314 +0530 IST m=+0.084274501
		// Consumed: 101 2026-01-02 17:47:51.2441413 +0530 IST m=+0.115484101
		// Consumer signaled at 2026-01-02 17:47:51.2441413 +0530 IST m=+0.115484101
		// Producer woke up at 2026-01-02 17:47:51.2447146 +0530 IST m=+0.116057401
		// Produced: 106 2026-01-02 17:47:51.2447146 +0530 IST m=+0.116057401
		// Producer signaled at 2026-01-02 17:47:51.2452299 +0530 IST m=+0.116572701
		// Producer is waiting at 2026-01-02 17:47:51.2593194 +0530 IST m=+0.130662101
		fmt.Println("Producer woke up at", time.Now())
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item, time.Now())
	b.cond.Signal() // Notify a waiting consumer
	fmt.Println("Producer signaled at", time.Now())
}

func (b *Buffer) Consume() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for len(b.items) == 0 {
		// When consumer is faster than producer, this condition will be true first.
		// Once the buffer is empty, the consumer has to wait for the producer to produce an item
		fmt.Println("Consumer is waiting at", time.Now())
		// This will release the lock and put the goroutine to sleep until signaled
		b.cond.Wait() // Wait until there is an item
		// 	Consumer is waiting at 2026-01-02 17:43:05.4478581 +0530 IST m=+0.001635001 // Due to cond.wait
		// Produced: 100 2026-01-02 17:43:05.4484517 +0530 IST m=+0.002228601
		// Producer signaled at 2026-01-02 17:43:05.4484517 +0530 IST m=+0.002228601 // Due to cond.Signal in Producer
		// Consumer woke up at 2026-01-02 17:43:05.4484517 +0530 IST m=+0.002228601
		// Consumed: 100 2026-01-02 17:43:05.4484517 +0530 IST m=+0.002228601
		fmt.Println("Consumer woke up at", time.Now())
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item, time.Now())
	b.cond.Signal() // Notify a waiting producer and release the waiting goroutine
	fmt.Println("Consumer signaled at", time.Now())
	return item
}

func producer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		b.Produce(i + 100)
		time.Sleep(1 * time.Millisecond) // Simulate production time
	}
}

func consumer(b *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		b.Consume()
		time.Sleep(100 * time.Millisecond) // Simulate consumption time
	}
}

func main() {

	buf := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(buf, &wg)
	go consumer(buf, &wg)

	wg.Wait()

}
