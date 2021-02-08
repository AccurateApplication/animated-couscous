package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

var textFlag = flag.String("text1", "", "text3")
var boolFlag = flag.Bool("bool", true, "bool123")

func main() {
	flag.Parse()
	fmt.Printf("string flagprint: %s\n", *textFlag)
	fmt.Printf("tool flag print: %t\n", *boolFlag)

	read := bufio.NewReader(os.Stdin)
	txt, err := read.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	x := sha256.Sum256([]byte("txt"))

	fmt.Printf("shasum text: %s\tgets:%x\n", txt, x)

	// fmt.Scanln(&sha)
}
