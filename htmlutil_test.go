package htmlutil

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const tmpl = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Page Title</title>
	</head>
	<body>
		<h1 class="greeting" id="title">Hello World!</h1>
		<p>2018</p>
	</body>
	</html>`

var doc, _ = html.Parse(strings.NewReader(tmpl))

func TestWalk(t *testing.T) {
	e := []string{"html", "head", "body", "meta", "title", "h1", "p"}
	m := make(map[string]bool)

	Walk(doc, func(n *html.Node) bool {
		m[n.Data] = true
		return false
	})

	for _, s := range e {
		if !m[s] {
			t.Errorf("Walk failed. element %s not walked", s)
		}
	}

	// should stop walking through the children
	m = make(map[string]bool)
	Walk(doc, func(n *html.Node) bool {
		m[n.Data] = true
		return IsElement(n, "body")
	})

	for _, s := range []string{"h1", "p"} {
		if m[s] {
			t.Error("Walk failed. body's children was walked")
		}
	}
}

func TestAttr(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "h1"); found {
			if cls := Attr(n, "class"); cls != "greeting" {
				t.Errorf("Attr failed. got class: %s, want: greeting", cls)
			}

			if id := Attr(n, "id"); id != "title" {
				t.Errorf("Attr failed. got id: %s, want: title", id)
			}
		}
		return
	})
}

func TestHasAttr(t *testing.T) {
	n := doc.LastChild // <html lang="en"></html>
	if !HasAttr(n, "lang", "en") {
		t.Error("HasAttr failed: node attr:", n.Attr)
	}
}

func TestHasText(t *testing.T) {
	var x *html.Node
	Walk(doc, func(n *html.Node) (found bool) {
		if found = HasText(n, "World"); found {
			x = n
		}
		return
	})

	if x.Type != html.TextNode || x.Data != "Hello World!" {
		t.Error("HasText failed. node type: %v, node data: %s", x.Type, x.Data)
	}
}

func TestIsElement(t *testing.T) {
	n := doc.LastChild
	if !IsElement(n, "html") {
		t.Errorf("IsElement failed. node type: %v, node data: %s", n.Type, n.Data)
	}
}

func TestText(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "title"); found {
			if txt := Text(n); txt != "Page Title" {
				t.Fatalf("Text failed. got: %s, want: Page Title", txt)
			}
		}
		return
	})
}

func TestInt(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "p"); found {
			i, err := Int(n)
			if err != nil {
				t.Fatal("Int failed:", err)
			}
			if i != 2018 {
				t.Errorf("Int failed. got %d, want: 2018", i)
			}
		}
		return
	})
}
