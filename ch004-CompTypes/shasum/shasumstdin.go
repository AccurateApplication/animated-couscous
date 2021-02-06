package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	txt, err := read.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	x := sha256.Sum256([]byte("txt"))

	fmt.Printf("shasum text: %s\tgets:%x\n", txt, x)

	// fmt.Scanln(&sha)
}
