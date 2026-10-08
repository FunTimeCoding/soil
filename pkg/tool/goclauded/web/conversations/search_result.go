package conversations

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func searchResult(
	c *conversation.Conversation,
	query string,
	kinds []string,
	open bool,
) gomponents.Node {
	name := c.Name

	if name == "" {
		name = constant.UnnamedSession
	}

	terms := strings.Fields(query)
	var snippets []gomponents.Node

	for _, h := range c.Hits {
		snippets = append(
			snippets,
			html.Div(
				html.Class("search-snippet"),
				extended.Get(hitsLocator(c.Session, query, kinds, h.Identifier)),
				extended.Target("#panel"),
				html.Small(gomponents.Text(blockLabel(h.Role, h.Kind))),
				gomponents.Group(highlighted(h.Snippet, terms)),
			),
		)
	}

	return html.Details(
		html.Class("search-result"),
		gomponents.If(open, html.Open()),
		html.Summary(
			html.Span(gomponents.Text(name)),
			html.Small(
				gomponents.Text(
					fmt.Sprintf(
						"%d hits · %s",
						c.Count,
						relativeTimestamp(c.Latest),
					),
				),
			),
		),
		gomponents.Group(snippets),
	)
}
