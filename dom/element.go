package dom

import (
	"strings"

	util "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

// An Element represents the HTML DOM element.
type Element interface {
	Children() []Element
	ClassList() []string
	ClassName() string
	FirstElementChild() Element
	GetAttribute(string) string
	GetElementsByClassName(string) []Element
	GetElementsByTagName(string) []Element
	InnerText() string
	ID() string
	LastElementChild() Element
	NextElementSibling() Element
	ParentElement() Element
	PreviousElementSibling() Element
	TagName() string
}

func (d dom) ClassList() []string {
	return strings.Fields(d.ClassName())
}

func (d dom) ClassName() string {
	return util.Attr(d.Node, "class")
}

func (d dom) FirstElementChild() Element {
	for n := d.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			return dom{n}
		}
	}
	return nil
}

func (d dom) GetAttribute(attr string) string {
	return util.Attr(d.Node, attr)
}

func (d dom) InnerText() string {
	var ss []string
	util.Walk(d.Node, func(n *html.Node) bool {
		if n.Type == html.TextNode {
			if s := strings.Trim(n.Data, "\n\t "); len(s) > 0 {
				ss = append(ss, n.Data)
			}
		}
		return util.IsElement(n, "script") || util.IsElement(n, "style")
	})
	return strings.Join(ss, "\n")
}

func (d dom) ID() string {
	return util.Attr(d.Node, "id")
}

func (d dom) LastElementChild() Element {
	for n := d.LastChild; n != nil; n = n.PrevSibling {
		if n.Type == html.ElementNode {
			return dom{n}
		}
	}
	return nil
}

func (d dom) NextElementSibling() Element {
	for n := d.NextSibling; n != nil; n = n.NextSibling {
		if n.Type == html.ElementNode {
			return dom{n}
		}
	}
	return nil
}

func (d dom) ParentElement() Element {
	if d.Parent != nil && d.Parent.Type == html.ElementNode {
		return dom{d.Parent}
	}
	return nil
}

func (d dom) PreviousElementSibling() Element {
	for n := d.PrevSibling; n != nil; n = n.PrevSibling {
		if n.Type == html.ElementNode {
			return dom{n}
		}
	}
	return nil
}

func (d dom) TagName() string {
	return strings.ToUpper(d.Data)
}
