package dom

import (
	"testing"
)

func TestTextContent(t *testing.T) {
	m := map[string]Node{
		"Coffee\nTea\nMilk\nCoke": doc.GetElementByID("drinks"),
		"List of Drinks":          doc.GetElementByID("main-2").FirstElementChild(),
		"// JavaScript content":   doc.GetElementsByTagName("script")[0],
		"":                        doc,
	}

	for want, node := range m {
		if got := node.TextContent(); got != want {
			t.Error("TextContent failed; node:", node)
			t.Log("got: ", got)
			t.Log("want:", want)
		}
	}
}
