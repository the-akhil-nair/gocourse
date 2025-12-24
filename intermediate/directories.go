package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

func main() {

	// err := os.Mkdir("subdir", 0755)
	// checkError(err)
	// checkError(os.Mkdir("subdir1", 0755))

	// Will remove directory and all its contents
	// defer os.RemoveAll("subdir1")

	// Since subdir1 already exists, we will use MkdirAll to avoid error.
	// os.WriteFile("subdir1/file", []byte(""), 0755)

	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		// In output the permissions are not shown as part of the Type()
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)
	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("../../.."))
	dir, err = os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// // filepath.Walk and filepath.WalkDir

	pathfile := "subdir"
	fmt.Println("Walking Directory")
	// WalkDir is more efficient than Walk and give better performance.
	// WalkDir uses os.DirEntry which avoids calling os.FileInfo on each file.
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println(path)
		return nil
	})

	checkError(err)

	defer func() {
		fmt.Println("Removing created directories")
		os.RemoveAll("subdir")
	}()
	// 	checkError(os.RemoveAll("subdir"))
	// checkError(os.Remove("subdir1"))
}
