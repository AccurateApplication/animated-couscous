package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "find links err: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}

	fmt.Println("vim-go")
}

//
func findLink(url string) ([]string, error) { // Multiple return values
	resp, err := http.Get(url) // os args in main
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOk { // if not ok, close
		resp.Body.Close()
		return nil, fmt.Errorf("getting url: %s : %s", url, resp.Status) // if err print error and url
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close() // doc = response body, then close resp
	if err != nil {
		return nil, fmt.Errorf("parse %s as html %v", url, err)
	}
	return visit(nil, doc), nil
}
