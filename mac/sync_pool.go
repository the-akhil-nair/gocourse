package mac

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var pool = sync.Pool{
		// New: func() any {
		// 	fmt.Println("Creating new Person instance")
		// 	return &Person{}
		// },
	}

	pool.Put(&Person{Name: "Alice", Age: 20})

	p1 := pool.Get()
	fmt.Printf("Person 1 (from pool): %+v\n", p1)

	// Get a Person instance from the pool
	// p1 := pool.Get().(*Person)
	// p1.Name = "Alice"
	// p1.Age = 30
	// fmt.Printf("Person 1: %+v\n", p1)

	pool.Put(p1) // Return the instance to the pool
	fmt.Println("Returned Person 1 to pool")

	// Get another Person instance from the pool
	p2 := pool.Get().(*Person)
	fmt.Printf("Person 2 (after getting from pool): %+v\n", p2)

	p3 := pool.Get()
	if p3 != nil {
		fmt.Println("Got Person3")
		p7 := p3.(*Person)
		p7.Name = "Bob"
		fmt.Printf("Person 3 (from pool): %+v\n", p7)
	} else {
		fmt.Println("Person3 is nil, pool is empty")
	}

	// p3 := pool.Get().(*Person) // This will create a new instance since pool is empty

	pool.Put(p2)
	pool.Put(p3)
	fmt.Println("Returned Person 2 and Person 3 to pool")

	p4 := pool.Get().(*Person)
	fmt.Printf("Person 4 (after getting from pool): %+v\n", p4)

	p5 := pool.Get().(*Person)
	fmt.Printf("Person 5 (after getting from pool): %+v\n", p5)
}
