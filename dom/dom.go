package dom

import (
	"strings"

	util "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

var (
	_ Element  = dom{}
	_ Document = dom{}
)

type dom struct {
	*html.Node
}

func (d dom) Children() []Element {
	var es []Element
	for c := d.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			es = append(es, dom{c})
		}
	}
	return es
}

func (d dom) GetElementsByClassName(names string) []Element {
	var (
		es   []Element
		list = strings.Fields(names)
	)
	if len(list) == 0 {
		return es
	}

	for c := d.FirstChild; c != nil; c = c.NextSibling {
		util.Walk(c, func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				if s := strings.Fields(util.Attr(n, "class")); len(s) != 0 {
					m := make(map[string]struct{}, len(s))
					for _, name := range s {
						m[name] = struct{}{}
					}

					for _, name := range list {
						if _, ok := m[name]; !ok {
							return false
						}
					}

					es = append(es, dom{n})
				}
			}
			return false
		})
	}
	return es
}

func (d dom) GetElementsByTagName(name string) []Element {
	var es []Element
	for c := d.FirstChild; c != nil; c = c.NextSibling {
		util.Walk(c, func(n *html.Node) bool {
			if util.IsElement(n, name) {
				es = append(es, dom{n})
			}
			return false
		})
	}
	return es
}

func (d dom) String() string {
	if d.Node.Type == html.DocumentNode {
		return "document"
	}

	s := d.Node.Data
	if id := d.ID(); id != "" {
		s += "#" + id
	}
	if cs := d.ClassList(); len(cs) != 0 {
		s = strings.Join(append([]string{s}, cs...), ".")
	}

	return s
}
