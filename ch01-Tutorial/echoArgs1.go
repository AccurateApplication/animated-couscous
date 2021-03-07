package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	var sep string
	for i := 0; i < len(os.Args); i++ {
		str += sep + os.Args[i]
		sep = " "
		fmt.Println(str)
	}
	// s := "" // this type of var may only be used inside a func, not pkg level var
}
