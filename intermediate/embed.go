package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed example.txt
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {
	// embed supports files and folders.
	// embeded files can not modified at runtime.
	// The content of the embedded file is available at compile time.
	// Consider using embed for static files like HTML templates, configuration files, etc.
	// Consider the size of the binary when embedding large files.
	fmt.Println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Embedded file content:", string(content))

	// fs.DirEntry provides a way to read the contents of a directory across multiple platforms.
	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

}
