// du written in go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

	// Print results
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++ // nfiles = number of files (?)
		nbytes += size
	}
	fmt.Println(nbytes) // ~1345 in cur dir
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("Seems to print 0 as size %d files %.1f GB\n", nfiles, float64(nbytes)/1e9) // nbytes/1e9 makes it human readable?
	// fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
	fmt.Println("test", nbytes/(1024*1024*1024))            // 0
	fmt.Printf("nfiles: %d\tnbytes: %v \n", nfiles, nbytes) // ~1345 cur dir

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
