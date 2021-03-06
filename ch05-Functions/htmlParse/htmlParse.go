package htmlParse

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func htmlParse() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse fail: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // append n data to stack (push)
		// https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// recursive, calls itself
		outline(stack, c) // gets copy of stack,
	}
}
