package intermediate

import (
	"errors"
	"fmt"
	"math"
)

// type interfaceError interface {
// 	PrintE() string
// }

type myError struct {
	message string
}

// Error is the interface method for error interface
// So myError now implements the error interface which means it can be used as an error type
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{message: "Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative number")
	}
	// compute the square root
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// Process data
	return nil
}

func main() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// data := []byte{}
	// // if err := process(data); err != nil {
	// err = process(data)

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	// --- error interface of builtin package
	err1 := eprocess()
	if err1 != nil {
		// Little bit more debugging gave me explanation that fmt.Println and fmt.Errorf will call the Error method explicitly.
		// These functions use reflection internally to:
		// * Check if the value implements the error interface.
		// * If it does, they call the Error() method on that value to get a string description, which is then printed to the console.
		fmt.Println(err1)
		fmt.Printf("Type of err1: %T\n", err1)
		return
	}

	//Available in buitin package.
	println("")

	// err := readData()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Data read successfully.")

}

// If this is used then the custom message looks like below:
// &{Custom error message}
// func eprocess() *myError {
// 	return &myError{message: "Custom error message"}
// }
