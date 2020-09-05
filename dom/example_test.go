package dom_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hunsin/go-htmlutil/dom"
)

// Example demonstrates a technique for getting the top 3 language views
// per day in Wikipedia.
func Example() {
	req, err := http.NewRequest(http.MethodGet, "https://www.wikipedia.org", nil)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := dom.Fetch(req)
	if err != nil {
		log.Fatalln(err)
	}

	top3 := doc.GetElementsByClassName("central-featured-lang")[:3]
	for _, elm := range top3 {
		lang := elm.GetElementsByTagName("strong")[0].TextContent()
		link := elm.FirstElementChild().GetAttribute("href")
		fmt.Println(lang, link)
	}
	// Output:
	// English //en.wikipedia.org/
	// 日本語 //ja.wikipedia.org/
	// Español //es.wikipedia.org/
}
