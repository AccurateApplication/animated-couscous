package main

import (
	"flag"
	"fmt"
	"strings"
)

// flags for cli program, can be displayd with -help when running program
var n = flag.Bool("n", false, "omit trail newline")
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse() // useful for CLI programs
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
