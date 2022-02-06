package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println("Begin.")
	a, err := listTextFiles(".")

	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("All text files: ", a)
}

func listTextFiles(path string) ([]string, error) {
	textFiles := []string{}

	x := func(pathInner string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".txt" {
			textFiles = append(textFiles, pathInner)
		}
		return nil
	}

	err := filepath.Walk(".", x)

	if err != nil {
		return nil, err
	}
	return textFiles, nil
}
