package main

import (
	"flag"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose")

func main() {
	var tick <-chan time.Time // receive channel
	if *verbose {
		tick = time.Tick(500 * time.Millisecond) // if verbose, send to tick channel
	}
	var nfiles, nbytes64 int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok { // if filesize closed
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

}
