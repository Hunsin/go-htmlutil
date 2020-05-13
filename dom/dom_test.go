package dom

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

const rawHTML = `
<!DOCTYPE html>
<html>
<head data-x="heading">
	<title>HTML Example</title>
	<link rel="stylesheet" href="style.css">
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body id="app">
	<style>
		/* style content */
	</style>

	<!-- comments are not displayed in the browser -->
	<nav id="navbar" class="flex" data-x="navbar">
		<a class="item nav-item nav-link" href="/">Home</a>
		<a class="item nav-item nav-link" href="/about">About</a>
		<a class="item nav-item nav-link" href="/careers">Careers</a>
		<div class="item nav-item right" id="menu"></div>
	</nav>

	<main>
		<article id="main-1">
			<h2>Article Title</h2>
			<p>This is a <i>paragraph</i>.</p>
		</article>
		<div id="main-2">
			<h2>List of Drinks</h2>
			<ul id="drinks" class="flex" data-x="drinks">
				<li class="item list-item">Coffee</li>
				<li class="item list-item">Tea</li>
				<li class="item list-item">Milk</li>
				<li class="item list-item">Coke</li>
			</ul>
		</div>
	</main>

	<form id="form">
		<input type="radio" name="chose" value="coffee">
		<label>Coffee</label>
		<input type="radio" name="chose" value="tea">
		<label>Tea</label>
		<input type="radio" name="chose" value="milk">
		<label>Milk</label>
		<input type="radio" name="chose" value="coke">
		<label>Coke</label>
		<input type="submit" value="Submit">
	</form>
	
	<script>
		// JavaScript content
	</script>
</body>
</html>
`

var doc Document

func TestMain(m *testing.M) {
	n, err := html.Parse(bytes.NewBufferString(rawHTML))
	if err != nil {
		panic(err)
	}
	doc = dom{n}

	os.Exit(m.Run())
}

func TestChildren(t *testing.T) {
	if children := doc.Children(); len(children) != 1 || children[0].TagName() != "HTML" {
		t.Error("Document.Children failed; got:", children)
	}

	var (
		main = doc.GetElementsByTagName("main")[0]
		want = []string{"article#main-1", "div#main-2"}
		got  = make([]string, 0, len(want))
	)
	for _, item := range main.Children() {
		got = append(got, fmt.Sprint(item))
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("Children failed; parent: ", main)
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestGetElementsByClassName(t *testing.T) {

	// test Document
	var (
		want = []string{"nav#navbar.flex", "ul#drinks.flex"}
		got  = make([]string, 0, len(want))
	)
	for _, item := range doc.GetElementsByClassName("flex") {
		got = append(got, fmt.Sprint(item))
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("GetElementsByClassName failed.")
		t.Log("got: ", got)
		t.Log("want:", want)
	}

	// test Element
	want = []string{"Coffee", "Tea", "Milk", "Coke"}
	got = make([]string, 0, len(want))
	for _, item := range doc.GetElementByID("drinks").GetElementsByClassName("list-item") {
		if fmt.Sprint(item) != "li.item.list-item" {
			t.Fatal("GetElementsByClassName failed; unexpected child element:", item)
		}
		got = append(got, item.InnerText())
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("GetElementsByClassName failed; unexpected innerText")
		t.Log("got: ", got)
		t.Log("want:", want)
	}

	// get elements with multiple class names
	for _, names := range []string{"nav-link", "item nav-link"} {
		for _, item := range doc.GetElementsByClassName(names) {
			if fmt.Sprint(item) != "a.item.nav-item.nav-link" {
				t.Error("GetElementsByClassName failed; unexpected element:", item)
			}
		}
	}
}

func TestGetElementsByTagName(t *testing.T) {

	// test Document
	main := doc.GetElementsByTagName("main")[0]
	if main.TagName() != "MAIN" {
		t.Errorf("GetElementsByTagName failed. got: %s, want: %s", main, "MAIN")
	}

	// test Element
	var (
		want = []string{"Article Title", "List of Drinks"}
		got  = make([]string, 0, len(want))
	)
	for _, h2 := range main.GetElementsByTagName("h2") {
		if h2.TagName() != "H2" {
			t.Fatalf("GetElementsByTagName failed; got: %s, want: %s", h2, "H2")
		}
		got = append(got, h2.InnerText())
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("GetElementsByTagName failed; unexpected innerText")
		t.Log("got: ", got)
		t.Log("want:", want)
	}
}

func TestString(t *testing.T) {
	m := map[string]interface{}{
		"document":          doc,
		"body#app":          doc.Body(),
		"nav#navbar.flex":   doc.GetElementByID("navbar"),
		"li.item.list-item": doc.GetElementByID("drinks").FirstElementChild(),
	}

	for want, element := range m {
		if got := fmt.Sprint(element); got != want {
			t.Errorf("String failed; got: %s, want: %s", got, want)
		}
	}
}
