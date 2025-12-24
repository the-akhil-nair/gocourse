package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Buffering: Buffering allow the movie to play smoothly without interruptions.
	// It stores a portion of the movie data in advance, so that if there are any delays in downloading,
	// the movie can continue to play from the buffered data without stopping.

	// Buffering reads and writes data in larger chunks, reducing the number of I/O operations.
	// This improves performance, especially for slow devices like disk drives or network connections.

	// strings.NewReader: Is a tool that helps us read data from a string as if it were a file or a stream.

	// NewReader: Creates a new buffered reader that reads from an underlying reader.
	// := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))

	// // Reading byte slice
	data := make([]byte, 20)
	// Read data from the buffered reader into the byte slice
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	// 	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
	fmt.Printf("Read %d bytes: %s\n", n, data)

	// Reading string will read until it encounters a newline character or delimiter in general.
	// Since reader already read some data, the next ReadString will continue from where it left off.
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read string:", line)

	// Since next time there is no newline character, it will read until the end of the input and encounter EOF and throws an error.
	// Since reader.ReadString is added the os.Stdin, it will wait for user input until Enter is pressed.
	// line, err = reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading string:", err)
	// 	return
	// }
	// fmt.Println("Read string:", line)

	// os.Stdout returns File struct that implements Write method which allows it to be used as an io.Writer.
	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	data = []byte("Hello, bufio package!\n")
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Added into internal buffer, not yet written to os.Stdout
	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
