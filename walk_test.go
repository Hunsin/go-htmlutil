package htmlutil

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var (
	elements = []string{"html", "head", "meta", "title", "body", "h1", "p"}
	walked   = make(map[string]bool)
)

func TestFirst(t *testing.T) {
	First(doc, func(n *html.Node) bool {
		walked[n.Data] = true
		return false
	})

	for _, e := range elements {
		if !walked[e] {
			t.Errorf("First failed. element %s not walked", e)
		}
	}

	// should stop walking through the children and siblings
	walked = make(map[string]bool)
	head := First(doc, func(n *html.Node) bool {
		walked[n.Data] = true
		return IsElement(n, "head")
	})

	if !IsElement(head, "head") {
		t.Errorf("First failed. got: %s, want: head", head.Data)
	}

	for _, e := range elements[2:] {
		if walked[e] {
			t.Errorf("First failed. head's children or sibling %s was walked", e)
		}
	}
}

func ExampleWalk() {
	s := `
	<div id="first">
		<div id="child"></div>
	</div>
	<div id="sibling"></div>`

	n, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}

	fn := func(n *html.Node) bool {
		found := IsElement(n, "div")
		if found {
			fmt.Println(Attr(n, "id"))
		}
		return found
	}

	Walk(n, fn)
	// Output:
	// first
	// sibling
}

func TestWalk(t *testing.T) {
	Walk(doc, func(n *html.Node) bool {
		walked[n.Data] = true
		return false
	})

	for _, e := range elements {
		if !walked[e] {
			t.Errorf("Walk failed. element %s not walked", e)
		}
	}

	// should stop walking through the children
	walked = make(map[string]bool)
	Walk(doc, func(n *html.Node) bool {
		walked[n.Data] = true
		return IsElement(n, "head")
	})

	for _, e := range elements[2:4] {
		if walked[e] {
			t.Errorf("Walk failed. head's children %s were walked", e)
		}
	}

	// should keep searching the siblings
	for _, e := range elements[4:] {
		if !walked[e] {
			t.Errorf("Walk failed. body or its children %s wasn't walked", e)
		}
	}
}
