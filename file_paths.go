package intermediate

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

	// Join paths using filepath.join
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	// filepath.Base and filepath.Dir will use file separator of the OS.
	fmt.Println(filepath.Base("/home/user/docs/"))
	fmt.Println(filepath.Dir("/home/user/docs/"))

	fmt.Println("Is relativePath variable absolute:", filepath.IsAbs(relativePath))
	//IsAbs gave false for absolute path
	// IsAbs checks if the path is absolute and a valid path else gave false.
	// 	absolutePath := "/home/user/docs/file.txt" gave fals because the path does not exist on Windows OS.
	// 	absolutePath := "D:\\learning\\gocourse\\data\\file.txt" works because its a valid absolute path on Windows.
	fmt.Println("Is absolutePath variable absolute:", filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)
	}
}
