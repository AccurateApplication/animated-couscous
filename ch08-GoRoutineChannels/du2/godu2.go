package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var nfiles, nbytes int64
var verbose = flag.Bool("v", false, "show verbose")

func main() {
	// Get initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."} // if args = 0 , use current dir
	}
	fileSizes := make(chan int64)

	go func() {
		for _, root := range roots { // roots either "." or files given to arg
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var tick <-chan time.Time // receive channel
	if *verbose {
		tick = time.Tick(500 * time.Millisecond) // if verbose, send to tick channel
	}

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok { // if filesize closed
				break loop
			}
			nfiles++
			nbytes += size // "ranges" over channel and sends size of dir to nbytes
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("Files: %v\tSize: %v (in bytes)\n", nfiles, nbytes)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name()) // "makes it a path"
			walkDir(subdir, fileSizes)                 // recursively walks subdir
		} else {
			fileSizes <- entry.Size() // entry = range(dir). calls func size()
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
