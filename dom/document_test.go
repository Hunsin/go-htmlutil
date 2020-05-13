package dom

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestBody(t *testing.T) {
	if body := doc.Body(); body == nil || body.TagName() != "BODY" || body.ID() != "app" {
		t.Error("Body failed. found:", body)
	}
}

func TestGetElementByID(t *testing.T) {
	for _, id := range []string{"navbar", "drinks"} {
		if elm := doc.GetElementByID(id); elm == nil || elm.GetAttribute("data-x") != id {
			t.Errorf("GetElementByID failed; id: %s, found: %s", id, elm)
		}
	}
}

func TestGetElementsByName(t *testing.T) {
	want := []string{"coffee", "tea", "milk", "coke"}
	got := make([]string, 0, len(want))

	for _, item := range doc.GetElementsByName("chose") {
		if item.TagName() != "INPUT" {
			t.Error("GetElementsByName failed. Got element tag name:", item.TagName())
		}
		got = append(got, item.GetAttribute("value"))
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("GetElementsByName failed; unexpected values")
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestHead(t *testing.T) {
	if head := doc.Head(); head == nil || head.GetAttribute("data-x") != "heading" {
		t.Error("Head failed; found:", head)
	}
}

func TestLinks(t *testing.T) {
	var (
		want = []string{"/", "/about", "/careers"}
		got  = make([]string, 0, len(want))
	)

	for _, anchor := range doc.Links() {
		if anchor.TagName() != "A" {
			t.Error("Links failed. Got element tag name:", anchor.TagName())
		}
		got = append(got, anchor.GetAttribute("href"))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Links failed; got: %s, want: %s", got, want)
	}
}

func TestTitle(t *testing.T) {
	if title := doc.Title(); title != "HTML Example" {
		t.Errorf("Title failed. got: %s, want: HTML Example", title)
	}
}

func TestFetch(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(rawHTML))
	})
	s := httptest.NewServer(h)
	r, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	d, err := Fetch(r)
	if err != nil {
		t.Fatal("Fetch failed: ", err)
	}

	if !reflect.DeepEqual(d, doc) {
		t.Error("Fetch failed; node structure not identical")
	}
}
