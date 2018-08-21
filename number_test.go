package htmlutil

import (
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

var p = First(doc, func(n *html.Node) bool {
	return IsElement(n, "p")
})

func numHelper(got, want interface{}, err error) error {
	if err != nil {
		return err
	}
	if got != want {
		return fmt.Errorf("got: %v, want: %v", got, want)
	}
	return nil
}

func TestInt(t *testing.T) {
	i, err := Int(p)
	if err = numHelper(i, -2019, err); err != nil {
		t.Error("Int failed.", err)
	}
}

func TestInt64(t *testing.T) {
	i, err := Int64(p)
	if err = numHelper(i, int64(-2019), err); err != nil {
		t.Error("Int64 failed.", err)
	}
}

func TestUint64(t *testing.T) {
	i, err := Uint64(p)
	if err = numHelper(i, uint64(2020), err); err != nil {
		t.Error("Uint64 failed.", err)
	}
}

func TestReal(t *testing.T) {

	// test float
	f, err := Real(p)
	if err = numHelper(f, float64(2018.8), err); err != nil {
		t.Error("Real failed.", err)
	}

	// test integer
	f, err = Real(p.FirstChild.NextSibling) // <b>-2019</b>
	if err = numHelper(f, float64(-2019), err); err != nil {
		t.Error("Real failed.", err)
	}
}
