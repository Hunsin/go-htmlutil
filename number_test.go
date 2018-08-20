package htmlutil

import (
	"testing"

	"golang.org/x/net/html"
)

func TestInt(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "p"); found {
			i, err := Int(n)
			if err != nil {
				t.Fatal("Int failed:", err)
			}
			if i != 2020 {
				t.Errorf("Int failed. got %d, want: 2020", i)
			}
		}
		return
	})
}

func TestInt64(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "p"); found {
			i, err := Int64(n)
			if err != nil {
				t.Fatal("Int64 failed:", err)
			}
			if i != 2020 {
				t.Errorf("Int failed. got %d, want: 2020", i)
			}
		}
		return
	})
}

func TestReal(t *testing.T) {
	Walk(doc, func(n *html.Node) (found bool) {
		if found = IsElement(n, "p"); found {

			// test float
			f, err := Real(n)
			if err != nil {
				t.Fatal("Real failed:", err)
			}
			if f != 2018.8 {
				t.Errorf("Real failed. got %f, want: 2018.8", f)
			}

			// test integer
			f, err = Real(n.LastChild)
			if err != nil {
				t.Fatal("Real failed:", err)
			}
			if f != 2020 {
				t.Errorf("Real failed. got %f, want: 2020", f)
			}
		}
		return
	})
}
