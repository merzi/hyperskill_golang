package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// write your code here
	fileList := make([]string, 1)
	if len(os.Args) <= 1 {
		fmt.Println("Directory is not specified")
		return
	}

	err := filepath.Walk(os.Args[1], func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fmt.Println(path)
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		return
	}
}
