package htmlLinkParser

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type TestCase struct {
	FileName string
	Expect   []Link
}

func Test_readFileIntoLinks(t *testing.T) {
	var tests = []TestCase{
		{"../examples/ex1.html", []Link{{
			Href: "/other-page",
			Text: "A link to another page",
		}}},
		{"../examples/ex2.html", []Link{{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		}, {
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on !",
		}}},
		{"../examples/ex3.html", []Link{{
			Href: "#",
			Text: "Login",
		}, {
			Href: "/lost",
			Text: "Lost? Need help?",
		}, {
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		}}},
		{"../examples/ex4.html", []Link{{
			Href: "/dog-cat",
			Text: "dog cat",
		}}},
	}

	for _, tc := range tests {
		t.Run(tc.FileName, func(t *testing.T) {
			links, err := readFileIntoLinks(tc.FileName)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(links, tc.Expect) {
				t.Errorf("Lists are not equal. \n\tExpected: %v \n\tGot: %v", tc.Expect, links)
			}
		})
	}

}
