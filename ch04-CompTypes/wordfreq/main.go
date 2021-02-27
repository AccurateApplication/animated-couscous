package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	dat, err := ioutil.ReadFile("./textfil")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(dat)

	s := string(dat[:])
	// s := strings.Trim(s, ".")
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	x := strings.Fields(s)
	// fmt.Println(x[:])
	//fmt.Println(string(dat[:]))
	m := make(map[string]int)

	for _, val := range x {
		// fmt.Println(val)
		m[val] += 1
		// fmt.Println("")
	}
	// fmt.Println(m)
	for key, val := range m {
		fmt.Println(val, key)
	}

}
