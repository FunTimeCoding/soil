package search_page

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func count(
	total int,
	collection string,
) gomponents.Node {
	noun := "results"

	if total == 1 {
		noun = "result"
	}

	text := fmt.Sprintf("%d %s", total, noun)

	if collection != "" {
		text = fmt.Sprintf("%s in %s", text, collection)
	}

	return html.P(gomponents.Text(text))
}
