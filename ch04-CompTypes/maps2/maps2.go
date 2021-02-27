package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // count unicode char
	var utflen [utf8.UTFMax + 1]int // UTF len()
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for { // while
		r, n, err := in.ReadRune() // returns: rune, nbytes, err
		if err == io.EOF {         // if eof/end , break
			break
		}
		if err != nil { // stnadard err check
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 { // if Char is unicode and nbytes = 1 {...
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("Rune\tCount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nLen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid utf8 characters\n", invalid)
	}

}

//maps each key to string

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

// Does it make any difference if func starts with capital letter?

var m = make(map[string]int)
