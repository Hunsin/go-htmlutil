package dom

import (
	"io"
	"net/http"

	util "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

type Document interface {
	Node

	Body() Element
	Children() []Element
	GetElementByID(id string) Element
	GetElementsByClassName(string) []Element
	GetElementsByName(string) []Element
	GetElementsByTagName(string) []Element
	Head() Element
	Links() []Element
	Title() string
}

func (d dom) Body() Element {
	return dom{util.First(d.Node, func(n *html.Node) bool {
		return util.IsElement(n, "body")
	})}
}

func (d dom) GetElementByID(id string) Element {
	if node := util.First(d.Node, func(n *html.Node) bool {
		return util.HasAttr(n, "id", id)
	}); node != nil {
		return dom{node}
	}
	return nil
}

func (d dom) GetElementsByName(name string) []Element {
	var es []Element
	util.Walk(d.Node, func(n *html.Node) bool {
		if n.Type == html.ElementNode && util.Attr(n, "name") == name {
			es = append(es, dom{n})
		}
		return false
	})
	return es
}

func (d dom) Head() Element {
	return dom{util.First(d.Node, func(n *html.Node) bool {
		return util.IsElement(n, "head")
	})}
}

func (d dom) Links() []Element {
	var es []Element
	util.Walk(d.Node, func(n *html.Node) bool {
		if util.IsElement(n, "a") && util.Attr(n, "href") != "" {
			es = append(es, dom{n})
		}
		return false
	})
	return es
}

func (d dom) Title() string {
	if title := util.First(d.Node, func(n *html.Node) bool {
		return util.IsElement(n, "title")
	}); title != nil {
		return util.Text(title)
	}
	return ""
}

func Fetch(req *http.Request) (Document, error) {
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	return Parse(r.Body)
}

func Parse(r io.Reader) (Document, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return dom{n}, nil
}
