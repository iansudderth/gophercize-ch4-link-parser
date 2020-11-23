package htmlLinkParser

import (
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func (l *Link) Equal(l2 *Link) bool {
	return l.Href == l2.Href && l.Text == l2.Text
}

func readFileIntoLinks(file string) ([]Link, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	var links []Link

	t, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			links = append(links, parseAnchor(node))
		} else {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(t)
	return links, nil
}

func parseAnchor(node *html.Node) Link {
	var text string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		}
	}

	var href string
	for _, d := range node.Attr {
		if d.Key == "href" {
			href = d.Val
			break
		}
	}

	return Link{
		Text: cleanString(text),
		Href: href,
	}
}

//var prefixSpace = regexp.MustCompile(`^\s+\S+`)
//var postfixSpace = regexp.MustCompile(`\S+\s+$`)

func cleanString(s string) string {
	trimmedString := strings.TrimSpace(s)
	//if prefixSpace.MatchString(s) {
	//	trimmedString = " " + trimmedString
	//}
	//if postfixSpace.MatchString(s) {
	//	trimmedString += " "
	//}
	return trimmedString

}
