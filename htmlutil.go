// Copyright 2018 Tsai, Hsiao-Chieh. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Package htmlutil provides some html utility functions.
package htmlutil

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Attr returns the value of keyed attribute under n. If no attribute
// matches the key, an empty string is returned.
func Attr(n *html.Node, key string) string {
	for i := range n.Attr {
		if n.Attr[i].Key == key {
			return n.Attr[i].Val
		}
	}
	return ""
}

// HasAttr returns whether n has any attribute which matches given
// key and val.
func HasAttr(n *html.Node, key, val string) bool {
	for i := range n.Attr {
		if n.Attr[i].Key == key && n.Attr[i].Val == val {
			return true
		}
	}
	return false
}

// HasText returns wether n is a html.TextNode and n.Data contains
// given sub string.
func HasText(n *html.Node, sub string) bool {
	return n.Type == html.TextNode && strings.Contains(n.Data, sub)
}

// IsElement returns whether n is a html.ElementNode and n.Data
// matches given data.
func IsElement(n *html.Node, data string) bool {
	return n.Type == html.ElementNode && n.Data == data
}

// Text returns the content of n's first text child node. If no
// html.TextNode was found, it returns empty string.
func Text(n *html.Node) string {
	var txt string
	First(n, func(n *html.Node) (found bool) {
		if found = (n.Type == html.TextNode); found {
			txt = n.Data
		}
		return
	})
	return txt
}

// Int returns the first integer it found in n and n's children.
// If no integer was found, it returns error.
func Int(n *html.Node) (int, error) {
	var i int
	var found bool
	var err error

	First(n, func(n *html.Node) bool {
		if n.Type == html.TextNode {
			if i, err = strconv.Atoi(n.Data); err == nil {
				found = true
				return true
			}
		}
		return false
	})

	if !found {
		return 0, errors.New("htmlutil: no number was found")
	}
	return i, nil
}

// Float64 returns the first number it found in n and n's children.
// If no number was found, it returns error.
func Float64(n *html.Node) (float64, error) {
	var f float64
	var found bool
	var err error

	First(n, func(n *html.Node) bool {
		if n.Type == html.TextNode {
			if f, err = strconv.ParseFloat(n.Data, 64); err == nil {
				found = true
				return true
			}
		}
		return false
	})

	if !found {
		return 0, errors.New("htmlutil: no number was found")
	}
	return f, nil
}
