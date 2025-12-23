package intermediate

import (
	"fmt"
)

type Speaker interface {
	Speak() string
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
}
