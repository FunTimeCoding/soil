package conversation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/block"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func Hits(
	blocks []*block.Block,
	matching map[string]bool,
	query string,
	around string,
) gomponents.Node {
	terms := strings.Fields(query)
	var result []gomponents.Node

	for _, b := range blocks {
		result = append(
			result,
			hitBlock(b, terms, matching[b.Identifier], b.Identifier == around),
		)
	}

	result = append(
		result,
		html.Div(html.ID("hit-counter"), html.Class("hit-counter")),
	)

	return gomponents.Group(result)
}
