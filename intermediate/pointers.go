package intermediate

import "fmt"

func main() {

	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)
	// No Pointer arithematic in go
	// fmt.Println(*ptr) // dereferencing a pointer
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	// Unsafe package to directly access memory
	// unsafe.Pointer(&x) converts x address to unsafe.Pointer

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
