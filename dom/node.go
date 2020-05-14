package dom

import (
	"strings"

	util "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

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

	// elementNode
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
