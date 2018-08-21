// Copyright 2018 Tsai, Hsiao-Chieh. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package htmlutil

import (
	"testing"

	"golang.org/x/net/html"
)

var (
	elements = []string{"html", "head", "meta", "title", "body", "h1", "p", "b", "span"}
	walked   = make(map[string]bool)
)

// checkWalked validates if elements are all walked by fn. If not, it
// calls t.Error.
func checkWalked(t *testing.T, fn string) {
	for _, e := range elements {
		if !walked[e] {
			t.Errorf("%s failed. element %s not walked", fn, e)
		}
	}
}

// recoder reallocates a new map to walked and returns a MatchFunc which
// records walked elements in the map. The stop should be a element name.
// If stop is specified, the func returns IsElement(n, stop).
func recoder(stop string) MatchFunc {
	walked = make(map[string]bool)

	return func(n *html.Node) bool {
		walked[n.Data] = true

		if stop != "" {
			return IsElement(n, stop)
		}
		return false
	}
}

func TestFirst(t *testing.T) {
	First(doc, recoder(""))
	checkWalked(t, "First")

	head := First(doc, recoder("head"))
	if !IsElement(head, "head") {
		t.Errorf("First failed. got: %s, want: head", head.Data)
	}

	// should stop walking through the children and siblings
	for _, e := range elements[2:] {
		if walked[e] {
			t.Errorf("First failed. head's children or sibling %s was walked", e)
		}
	}
}

func TestLast(t *testing.T) {
	Last(doc, recoder(""))
	checkWalked(t, "Last")

	span := Last(doc, recoder("span"))
	if !IsElement(span, "span") {
		t.Errorf("Last failed. got: %s, want: span", span.Data)
	}

	for _, e := range []string{"head", "meta", "title", "h1", "2019"} {
		if walked[e] {
			t.Errorf("Last failed. Element %s was walked", e)
		}
	}
}

func TestWalk(t *testing.T) {
	Walk(doc, recoder(""))
	checkWalked(t, "Walk")

	// should stop walking through the head's children
	Walk(doc, recoder("head"))
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
