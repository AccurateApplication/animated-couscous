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
	var stack []string
	stack = append(stack, "test")
	top := stack[len(stack)-1]
	fmt.Println(stack, top)

	p1 := [5]int{1, 2, 3, 4, 5}
	test := len(p1)
	fmt.Println(test)
	fmt.Println(p1, &p1)
	revArray(&p1)
	fmt.Println(p1, &p1)

}

// makes unnecessary copy of array to slice
func revArray(ptrArr *[5]int) {
	var temp []int
	fmt.Println(temp)
	for i := len(ptrArr) - 1; i >= 0; i-- {
		fmt.Println(ptrArr[i], i)
		// temp = append(*temp, ptrArr[i])
		temp = append(temp, ptrArr[i])

	}
	fmt.Println("ching", temp)
	copy(ptrArr[:], temp)
	//temp[

	//x := 0
	//for i := 4; x < len(ptrArr)-1; {
	//	// ptrArr[x] = ptrArr[i]
	//	(*ptrArr)[x] = (ptrArr)[i]
	//	(*ptrArr)[i] = (ptrArr)[x]
	//	x--
	//	i++
	//}
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

// for i, x := 0, len(s)-1; i < x; i, x = i+1, x-1 {
// 	s[i], s[x] = s[x], s[i]
// }

func revSlice(s []int) { // s = 10
	for i, x := 0, len(s)-1; i < x; i, x = i+1, x-1 {
		s[i], s[x] = s[x], s[i]
	}
}
