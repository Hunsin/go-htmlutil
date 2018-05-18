// Copyright 2018 Tsai, Hsiao-Chieh. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package htmlutil_test

import (
	"fmt"
	"strings"

	"github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

func ExampleAttr() {
	frag := `<html lang="en"></html>`
	n, _ := html.Parse(strings.NewReader(frag)) // ignore error

	lang := htmlutil.Attr(n.FirstChild, "lang")
	fmt.Println(lang)
	// Output: en
}

func ExampleIsElement() {
	frag := "<html></html>"
	n, _ := html.Parse(strings.NewReader(frag)) // ignore error

	b := htmlutil.IsElement(n.FirstChild, "html")
	fmt.Println(b)
	// Output: true
}

func ExampleText() {
	frag := "<p>Hello <strong>World</strong></p>"
	n, _ := html.Parse(strings.NewReader(frag)) // ignore error

	txt := htmlutil.Text(n)
	fmt.Println(txt)
	// Output: Hello
}

func ExampleInt() {
	frag := "<p>2018</p>"
	n, _ := html.Parse(strings.NewReader(frag)) // ignore error

	i, _ := htmlutil.Int(n)
	fmt.Println(i)
	// Output: 2018
}

func ExampleWalk() {
	frag := `
	  <div id="first">
	    <div id="child"></div>
	  </div>
	  <div id="sibling"></div>`
	n, err := html.Parse(strings.NewReader(frag))
	if err != nil {
		fmt.Println(err)
	}

	fn := func(n *html.Node) bool {
		found := htmlutil.IsElement(n, "div")
		if found {
			fmt.Println(htmlutil.Attr(n, "id"))
		}
		return found
	}

	htmlutil.Walk(n, fn)
	// Output:
	// first
	// sibling
}
