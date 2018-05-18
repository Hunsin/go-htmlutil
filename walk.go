// Copyright 2018 Tsai, Hsiao-Chieh. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package htmlutil

import "golang.org/x/net/html"

// A MatchFunc is a function which returns if a html node matches
// certain condition.
type MatchFunc func(*html.Node) bool

// First works like Walk(n, fn) but returns the node once fn(node)
// returns true.
//
// To compare First(), Last() and Walk(), consider a MatchFunc,
// fn looks like this:
// 	func(n *html.Node) bool {
// 	    return htmlutil.IsElement(n, "li")
// 	}
//
// Base on following HTML fragment, they work as:
// 	<ul>             | First(n, fn) | Last(n, fn) | Walk(n, fn) |
//	  <li>...</li>   | return node  | not walked  | walked      |
//	  <li>...</li>   | not walked   | not walked  | walked      |
//	  <li>...</li>   | not walked   | return node | walked      |
// 	</ul>
func First(n *html.Node, fn MatchFunc) *html.Node {
	if fn(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := First(c, fn); found != nil {
			return found
		}
	}
	return nil
}

// Last works like First(n, fn) but searches nodes in different direction.
// It returns the last matched node in the node tree.
func Last(n *html.Node, fn MatchFunc) *html.Node {
	if fn(n) {
		return n
	}

	for c := n.LastChild; c != nil; c = c.PrevSibling {
		if found := Last(c, fn); found != nil {
			return found
		}
	}
	return nil
}

// Walk calls fn with n and all the children under n. If fn(n)
// returns true, it keeps searching the node's siblings but children.
func Walk(n *html.Node, fn MatchFunc) {
	if fn(n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Walk(c, fn)
	}
}
