package htmlLinkParser

import (
	"golang.org/x/net/html"
	"os"
)

type Link struct {
	Href string
	Text string
}

func readFileIntoTokens(file string) ([]*html.Node, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var nodes []*html.Node

	t, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode {
			nodes = append(nodes, node)
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(t)
	return nodes, nil
}
