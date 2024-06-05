package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visitAllFiles() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

/* Walks recursively through the current file path
 * For each file in the path, print out the size of the file
 */
func betterVisitAllFiles() {
	args := os.Args[1:]

	err := filepath.WalkDir(args[0],
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			i, err := d.Info()
			if err != nil {
				return err
			}

			fmt.Println(path, i.Size())
			return nil
		})

	if err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("First")
	visitAllFiles()

	fmt.Println("Second")
	betterVisitAllFiles()
}
