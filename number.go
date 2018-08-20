package htmlutil

import (
	"fmt"
	"strconv"

	"golang.org/x/net/html"
)

// number finds the first text node which the data is convertible by fn.
// The numType indicates that which type the fn is intend to parse.
func number(n *html.Node, fn func(string) error, numType string) error {
	var found bool

	First(n, func(n *html.Node) bool {
		found = (n.Type == html.TextNode && fn(n.Data) == nil)
		return found
	})

	if !found {
		return fmt.Errorf("htmlutil: no %s was found", numType)
	}
	return nil
}

// Int returns the first integer it found in n and n's children.
// Floating-point numbers are ignored. If no integer was found,
// it returns error.
func Int(n *html.Node) (int, error) {
	var i int

	return i, number(n, func(s string) (err error) {
		i, err = strconv.Atoi(s)
		return err
	}, "integer")
}

// Int64 acts same as Int(n) but returns an interger of type int64.
func Int64(n *html.Node) (int64, error) {
	var i int64

	return i, number(n, func(s string) (err error) {
		i, err = strconv.ParseInt(s, 10, 64)
		return err
	}, "integer")
}

// Real returns the first number it found in n and n's children.
// If no number was found, it returns error.
func Real(n *html.Node) (float64, error) {
	var f float64

	return f, number(n, func(s string) (err error) {
		f, err = strconv.ParseFloat(s, 64)
		return err
	}, "number")
}
