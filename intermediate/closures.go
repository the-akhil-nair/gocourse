package intermediate

import "fmt"

func main() {

	// Closures are function which gives way to save the state and behaviour.
	// sequence := adder()

	// In subsequent calls the anonymous function is called and not the adder function.
	// sequence is called closure function.
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// In this case adder function is called and i value is reset.
	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtracter := func() func(int) int {

		countdown := 99
		return func(x int) int {
			// The value of countdown will be kept in the memory for closure function.
			countdown -= x
			return countdown
		}
	}()

	// Using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	// anonymous function
	return func() int {
		// i value is set into state and i value is increasing with every run
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
