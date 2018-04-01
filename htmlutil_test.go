// Copyright 2018 Tsai, Hsiao-Chieh. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

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
		<p>2018<span>2019</span></p>
	</body>
	</html>`

var doc, _ = html.Parse(strings.NewReader(tmpl))

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
		t.Errorf("HasText failed. node type: %v, node data: %s", x.Type, x.Data)
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
