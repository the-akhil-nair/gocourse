package intermediate

import (
	"fmt"
	"math/rand"
)

func main() {
	// Pseudo random numbers are generated using algorithms rather than from truly random physical processes.
	// The numbers generated are determined by an initial value known as a seed.
	// If you use the same seed value, you will get the same sequence of random numbers each time you run your program.
	// Random number generator with a fixed seed
	// val := rand.New(rand.NewSource(42))

	// randonm is not useful in concurrent programs as multiple goroutines may access the same source
	// Also use cryptorand package for cryptographic purposes

	// We dont have fixed seed now.
	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// This will give number between 5 to 10 exclusive
	// fmt.Println(rand.Intn(6) + 5)
	// Random package seeding automatically with current time
	// starting from 0 to 100 inclusive.
	// fmt.Println(rand.Intn(101))
	// fmt.Println(val.Intn(101))
	// fmt.Println(rand.Float64()) // between 0.0 and 1.0

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		// This is done as we are not including the 0.
		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show the results
		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user wants to roll again
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}
