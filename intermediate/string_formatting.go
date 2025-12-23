package intermediate

import "fmt"

func main() {

	num := 1
	// Padding 0s to the left to make it 5 digits
	fmt.Printf("%05d\n", num)

	message := "Hello"
	// The fixed width is 10 characters and if the string is shorter, it will be padded with spaces.
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	// If the string exceeds the specified width, it will be printed in full without truncation.
	message3 := "Hello, World!. Welcome to Go string formatting."
	fmt.Printf("|%10s|\n", message3)
	fmt.Printf("|%-10s|\n", message3)

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)
	fmt.Println(message2)

	// sqlQuery := `SELECT * FROM users WHERE age > 30`

}
