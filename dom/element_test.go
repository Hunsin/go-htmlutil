package dom

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClassList(t *testing.T) {
	a := doc.GetElementByID("navbar").FirstElementChild()
	want := []string{"item", "nav-item", "nav-link"}
	if got := a.ClassList(); !reflect.DeepEqual(got, want) {
		t.Error("ClassList failed. Element:", a)
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestClassName(t *testing.T) {
	a := doc.GetElementByID("navbar").FirstElementChild()
	want := "item nav-item nav-link"
	if got := a.ClassName(); got != want {
		t.Error("ClassName failed. Element:", a)
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestFirstElementChild(t *testing.T) {
	elm := doc.GetElementsByTagName("main")[0].FirstElementChild()
	want := "article#main-1"
	if fmt.Sprint(elm) != want {
		t.Errorf("FirstElementChild failed; got: %s, want: %s", elm, want)
	}

	elm = elm.FirstElementChild()
	want = "Article Title"
	if txt := elm.InnerText(); txt != want {
		t.Errorf("FirstElementChild failed; got innerText: %s, want: %s", txt, want)
	}
}

func TestInnerText(t *testing.T) {
	want := "Coffee\nTea\nMilk\nCoke"
	if got := doc.GetElementByID("drinks").InnerText(); got != want {
		t.Error("InnerText failed")
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestID(t *testing.T) {
	if got := doc.GetElementsByTagName("nav")[0].ID(); got != "navbar" {
		t.Errorf("ID failed; got: %s, want: %s", got, "navbar")
	}

	if got := doc.GetElementsByTagName("article")[0].ID(); got != "main-1" {
		t.Errorf("ID failed; got: %s, want: %s", got, "main-1")
	}
}

func TestLastElementChild(t *testing.T) {
	menu := doc.GetElementByID("navbar").LastElementChild()
	want := "div#menu.item.nav-item.right"
	if fmt.Sprint(menu) != want {
		t.Errorf("LastElementChild failed; got ID: %s, want: %s", menu, want)
	}

	main := doc.GetElementsByTagName("main")[0]
	want = "Coke"
	if txt := main.LastElementChild().LastElementChild().LastElementChild().InnerText(); txt != want {
		t.Errorf("LastElementChild failed; got innerText: %s, want: %s", txt, want)
	}
}

func TestNextElementSibling(t *testing.T) {
	li := doc.GetElementByID("drinks").FirstElementChild()
	txt := []string{"Coffee", "Tea", "Milk", "Coke"}
	for ; li != nil; li = li.NextElementSibling() {
		if li.TagName() != "LI" || li.InnerText() != txt[0] {
			t.Error("NextElementSibling failed")
		}
		txt = txt[1:]
	}
}

func TestParentElement(t *testing.T) {
	body := doc.GetElementByID("app")
	if parent := body.ParentElement(); parent.TagName() != "HTML" {
		t.Error("ParentElement failed")
		t.Log("element:   ", body)
		t.Log("got parent:", parent)
	}

	ul := doc.GetElementByID("drinks")
	if parent := ul.ParentElement(); fmt.Sprint(parent) != "div#main-2" {
		t.Error("ParentElement failed")
		t.Log("element:   ", ul)
		t.Log("got parent:", parent)
	}
}

func TestPreviousElementSibling(t *testing.T) {
	li := doc.GetElementByID("drinks").LastElementChild()
	txt := []string{"Coke", "Milk", "Tea", "Coffee"}
	for ; li != nil; li = li.PreviousElementSibling() {
		if li.TagName() != "LI" || li.InnerText() != txt[0] {
			t.Error("NextElementSibling failed")
		}
		txt = txt[1:]
	}
}
