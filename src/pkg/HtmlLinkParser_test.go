package htmlLinkParser

import (
	"testing"
)

func Test_readFileIntoTokens(t *testing.T) {
	nodes, err := readFileIntoTokens("../examples/ex1.html")
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) < 1 {
		t.Error("Length of Nodes should be greater than 1")
	}
}
