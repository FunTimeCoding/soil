package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
)

func Print(
	n *chromedp.Node,
	attribute []string,
) {
	console.Format("  XPath: %s\n", n.FullXPath())

	for _, a := range attribute {
		console.Format("  %s: %s\n", a, n.AttributeValue(a))
	}
}
