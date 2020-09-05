package dom

import (
	"strings"

	util "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

// A Node represents the DOM Node interface.
type Node interface {
	TextContent() string
}

func (d dom) TextContent() string {
	switch d.Node.Type {
	case html.CommentNode, html.TextNode:
		return d.Data

	case html.DoctypeNode, html.DocumentNode:
		return ""
	}
	// else html.ElementNode

	var ss []string
	util.Walk(d.Node, func(n *html.Node) bool {
		if n.Type == html.TextNode {
			if s := strings.Trim(n.Data, "\n\t"); len(s) > 0 {
				ss = append(ss, s)
			}
		}
		return false
	})
	return strings.Join(ss, "\n")
}
