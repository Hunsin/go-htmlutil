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
// Consider a MatchFunc, fn looks like this:
// 	func(n *html.Node) bool {
// 		return IsElement("li")
// 	}
//
// And a HTML fragment looks like this:
// 	<ul>
//	  <li>...</li>	<-- First(n, fn) stops here
//	  <li>...</li>
//	  <li>...</li>	<-- Walk(n, fn) loops all the <li></li> elements
// 	</ul>
//
// Both First(n, fn) and Walk(n, fn) won't loop the children under <li></li>
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
