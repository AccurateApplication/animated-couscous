package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	//"os"
)

var rmdirs []func()

func main() {
	numbers := sum(5, 3, 4, 35)
	werds := joinStrengs("hello", "waa", "quuu", "chee\n")
	fmt.Println(numbers, "\n\n", werds)
	werds2 := []byte(werds)
	writeStuffFile("temp", werds2)
	x := echoFile("temp")
	fmt.Printf("echo file: %s", x)

}
func sum(num ...int) int {
	total := 0
	for _, val := range num {
		total += val
	}
	return total
}

// joins words
func joinStrengs(word ...string) string {
	var finalWord string
	finalWord = strings.Join(word, finalWord)
	//for _, str := range word {
	//	strings.Join(str, finalWord) // dont work, str join want []string
	//}
	return finalWord
}

func writeStuffFile(filename string, textWrite []byte) {
	err := ioutil.WriteFile(filename, textWrite, 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func echoFile(file string) string {
	data, err := ioutil.ReadFile(file) // data = byte (string byte?)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(data))
	content := string(data)
	return content
}
