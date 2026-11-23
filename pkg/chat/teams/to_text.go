package teams

import (
	"golang.org/x/net/html"
	"strings"
)

func toText(markup string) string {
	n, e := html.Parse(strings.NewReader(markup))

	if e != nil {
		return markup
	}

	var b strings.Builder
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		switch {
		case n.Type == html.TextNode:
			b.WriteString(n.Data)
		case n.Type == html.ElementNode && n.Data == "br":
			b.WriteString("\n")
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}

		if n.Type == html.ElementNode && (n.Data == "p" || n.Data == "div") {
			b.WriteString("\n")
		}
	}
	walk(n)

	return strings.TrimSpace(b.String())
}
