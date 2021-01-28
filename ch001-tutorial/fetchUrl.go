package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] { // all args except calling arg = url
		// add if strings.hasprefix https:// = false then insert before url
		testString := strings.HasPrefix(url, "https://")
		fmt.Printf("test w/ url: %s\tw/ result: %t\n", url, testString)
		httpsString := "https://"
		if testString == false {
			fmt.Printf("url before: %s\n", url)
			url = httpsString + url
			fmt.Printf("fixed url: %s\n\n", url)
		}
		// If false, insert https before url, somehow
		// fmt.Printf(strings.HasPrefix(url, "https://"))
		resp, err := http.Get(url) // http request
		fmt.Println("http status", resp.StatusCode, url, resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1) // exit status 1 i suppose?
		}

		// b, err := ioutil.ReadAll(resp.Body) // b is response body? yes
		dst := os.Stdout
		testCopy, err := io.Copy(dst, resp.Body)
		fmt.Println("testCopy now used", testCopy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading: %s\t%v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("hmm\t %s\t testCopy: %s\t\n", dst, testCopy)
		resp.Body.Close()
		// fmt.Printf("%s\n", b)
	}
	str1 := "google.com"
	str2 := "https://"
	newString := str2 + str1
	fmt.Printf("str1: %s\tstr2: %s\tnew string: %s\n", str1, str2, newString)
	str1 = str2 + str1
	fmt.Printf("str 1 %s\n", str1)
}
