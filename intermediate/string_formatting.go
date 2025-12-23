package intermediate

import "fmt"

func main() {

	num := 1
	// Padding 0s to the left to make it 5 digits
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)
	fmt.Println(message2)

	// sqlQuery := `SELECT * FROM users WHERE age > 30`

}
