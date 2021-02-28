package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Statuscode != http.StatusOK {
		return nil, fmt.Errorf("Get: %s : %s", url, resp.Status)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" { // if elementnode and html node data = "a"
			for _, a := range n.Attr {
				if a.Key != "href" { // if not "href" discard and continue loop
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val) // ??
				if err != nil {                            // bad urls
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil) // ????????
	return links, nil
}
