package intermediate

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Vehicle interface {
	Wheels() int
	IsHeavy() bool
}

type Car struct{}

type Truck struct{}

type Motorcycle struct{}

func (c Car) Wheels() int {
	return 4
}

func (c Car) IsHeavy() bool {
	return false
}

func (t Truck) Wheels() int {
	return 6
}

func (t Truck) IsHeavy() bool {
	return true
}

func (m Motorcycle) Wheels() int {
	return 2
}

func (m Motorcycle) IsHeavy() bool {
	return false
}

func truckVehicle() Vehicle {
	return &Truck{}
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	var s Speaker
	d := Dog{}
	/* s is a variable of type Speaker, but since it is declared without assignment, its value is nil. d is a struct instance
	of type Dog. When comparing an interface with a concrete type, Go first checks whether the interface variable (s) is holding
	a value of the same concrete type (Dog). In this case, s is still nil, meaning it does not hold any Dog value. d, on the
	other hand, is an actual struct instance of Dog. Since s does not hold d, the comparison s == d evaluates to false.
	*/
	fmt.Println(s == d)
	// Now in this case we are assigning the Dog instance to the Speaker interface variable.
	s = Speaker(d)
	fmt.Println(s == d)
	fmt.Println(s.Speak())

	t := truckVehicle()
	fmt.Printf("Truck has %d wheels. Is it heavy? %t\n", t.Wheels(), t.IsHeavy())
	fmt.Printf("Truck %#v\n", t)
}
