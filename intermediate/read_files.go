package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()

	fmt.Println("File opened successfully.")

	// // Read the contents of the opened file
	// data := make([]byte, 1024) // Buffer to read data into
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading data from file:", err)
	// 	return
	// }

	// fmt.Println("File content:", string(data))
	// Difference between bufio.Reader and bufio.Scanner is that Scanner provides a convenient
	// interface for reading data such as a file line by line or token by token whereas Reader provides more general methods
	// for reading data in various ways (byte slices, strings, etc.).
	scanner := bufio.NewScanner(file)

	// Read line by line
	// Since last line is shown as EOF without newline character, it will not be printed unless we add an extra newline character at the end of the file.
	// Scan is something like next() in other languages.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

}
