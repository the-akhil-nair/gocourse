package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempFileName := "temporaryFile"

	// If the first parameter is an empty string, the file will be created in the default temporary-files directory
	// of the operating system.
	// Also add the prefix to the file name.
	// tempFile, err := os.CreateTemp("", tempFileName)
	// checkError(err)

	// fmt.Println("Temporary file created:", tempFile.Name())

	// defer os.Remove(tempFile.Name())
	// defer tempFile.Close()

	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created:", tempDir)
}
