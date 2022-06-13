package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var sorting = 0
	// write your code here
	if len(os.Args) <= 1 {
		fmt.Println("Directory is not specified")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter file format:")
	scanner.Scan()
	fileType := scanner.Text()

	fmt.Println("Size sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")
	for scanner.Scan() {
		sorting, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Wrong option")
			continue
		}

		if sorting == 1 || sorting == 2 {
			break
		}
		fmt.Println("Wrong option")
	}

	fList := Files{rootPath: os.Args[1]}
	if len(strings.TrimSpace(fileType)) > 1 {
		fList.determineFilesByExtension(strings.TrimSpace(fileType))
	} else {
		fList.determineAllFiles()
	}

	fList.printSortedFiles(sorting)
}

type Files struct {
	rootPath string
	files    map[int64][]File
}

func (f Files) setRootPath(rootPath string) {
	f.rootPath = rootPath
}

func (f Files) determineAllFiles() {
	f.determineFilesByExtension("")
}

func (f Files) determineFilesByExtension(extension string) {
	f.files = make(map[int64][]File)
	err := filepath.Walk(os.Args[1], func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if len(strings.TrimSpace(extension)) > 0 && strings.TrimSpace(extension) != filepath.Ext(path) {
			return nil
		}
		fmt.Println(info.Name())
		if _, ok := f.files[info.Size()]; ok {
			f.files[info.Size()] = append(f.files[info.Size()],
				File{path: path, size: info.Size(), extension: filepath.Ext(path)})
		} else {
			f.files[info.Size()] = []File{{path: path, size: info.Size(), extension: filepath.Ext(path)}}
		}

		return nil
	})
	if err != nil {
		return
	}
}

func (f Files) getFiles(sorting int) map[int64][]File {
	fileSizes := make([]int64, 0, len(f.files))
	fmt.Println(f.files)
	for k, v := range f.files {
		fmt.Println(k)
		fmt.Println(v)
		fileSizes = append(fileSizes, k)
	}

	if sorting == 1 {
		sort.Slice(fileSizes, func(i int, j int) bool {
			return fileSizes[i] > fileSizes[j]
		})
	} else {
		sort.Slice(fileSizes, func(i, j int) bool {
			return fileSizes[i] < fileSizes[j]
		})
	}

	output := make(map[int64][]File)
	for _, entry := range fileSizes {
		output[entry] = f.files[entry]
	}

	return output
}

func (f Files) printSortedFiles(sorting int) {
	files := f.getFiles(sorting)

	for size, fileSlice := range files {
		fmt.Printf("%d bytes\n", size)
		for _, file := range fileSlice {
			fmt.Println(file.path)
		}
		fmt.Println()
	}
}

type File struct {
	path      string
	extension string
	size      int64
}
