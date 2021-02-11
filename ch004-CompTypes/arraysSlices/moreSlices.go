package main

import (
	"fmt"
)

func main() {

	strSlice := []string{"hej1", "hej2", "", "hej3", "", ""}
	fmt.Println(strSlice)         // hej1 hej2 hej3
	strSlice = nonEmpty(strSlice) // if we just call function on strSlice, will duplicate values
	fmt.Println(strSlice)         // hej1 hej2 hej3
	strSlice2 := []string{"hej1", "hej2", "", "hej3", "", ""}
	strSlice2 = nonEmpty2(strSlice2)
	// fmt.Println(strSlice2)

}

func nonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// we can also use append
func nonEmpty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
