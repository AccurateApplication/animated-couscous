// du written in go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name()) // filepath join..?
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size() // entry size?
		}
	}
}

// dirent return entries of directory (dir)
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
