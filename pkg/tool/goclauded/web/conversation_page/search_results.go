package conversation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SearchResults(
	found []*conversation.Conversation,
	query string,
	kinds []string,
	indexed int64,
	total int64,
) gomponents.Node {
	var result []gomponents.Node

	if indexed < total {
		result = append(
			result,
			html.Small(
				html.Class("search-progress"),
				gomponents.Text(
					fmt.Sprintf("indexing %d/%d conversations", indexed, total),
				),
			),
		)
	}

	if len(found) == 0 {
		result = append(
			result,
			html.Small(
				html.Class("search-progress"),
				gomponents.Text("No conversation holds every term."),
			),
		)
	}

	for i, c := range found {
		result = append(result, searchResult(c, query, kinds, i == 0))
	}

	return gomponents.Group(result)
}
