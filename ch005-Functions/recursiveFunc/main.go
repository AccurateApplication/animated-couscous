package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
recursive functions may sometimes call themselves, either directly or indirectly

*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find link: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

// links found in n append to links slice
func visit(links []string, n *html.Node) []string { // result = links slice

	if n.Type == html.ElementNode && n.Data == "a" { //
		for _, a := range n.Attr { // loop over html node attributes
			if a.Key == "href" {
				links = append(links, a.Val) // if <a href>, append value of it to links slice
			}
		}

		// n.FirstChild = struct in html package, struct below
		// ( type Node struct{ ... FirstChild,NextSibling ---  *Node} )
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			// Recursive part. visit calls on itself for each of n.Children (?), which is held in FirstChild list
			links = visit(links, c)
		}
	}
	return links
}
