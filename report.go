package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

type File struct {
	Path string
	Size int64
}

var files []File

func handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	files = append(files, File{path, info.Size()})

	return nil
}

func displaySize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	}
	if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	}
	return fmt.Sprintf("%.2f GB", float64(size)/1024/1024/1024)
}

func main() {
	if len(os.Args) > 2 {
		log.Fatal("\nUsage: go run utility.go <directory>\n<directory>: optional, default value is current directory")
	}
	root := "./"
	if len(os.Args) == 2 {
		root = os.Args[1]
	}

	fmt.Println("Listing files in ", root)
	fmt.Println()
	err := filepath.Walk(root, handleFile)
	if err != nil {
		log.Println(err)
	}

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Size < files[j].Size
	})

	totalSize := int64(0)
	for index, file := range files {
		totalSize += file.Size
		normalizedIndex := len(files) - index
		step := math.Min(33, float64(len(files))/3)
		if normalizedIndex <= 100 {
			fileName := filepath.Base(file.Path)
			output := "• " + fileName + " " + displaySize(file.Size)
			if float64(normalizedIndex)/step < 1 {
				color.Red(output)
			} else if float64(normalizedIndex)/step < 2 {
				color.Yellow(output)
			} else {
				color.Green(output)
			}
		}
	}
	fmt.Println("\nTotal size: ", displaySize(totalSize))
}
