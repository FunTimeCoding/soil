package conversation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"github.com/funtimecoding/soil/pkg/web/component/highlight"
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
		locator := hitsLocator(c.Session, query, kinds, h.Identifier)
		snippets = append(
			snippets,
			html.A(
				html.Class("search-snippet"),
				html.Href(locator),
				extended.Get(locator),
				extended.Target(constant.PanelTarget),
				extended.Push("true"),
				html.Small(gomponents.Text(blockLabel(h.Role, h.Kind))),
				highlight.New(h.Snippet, terms...).Node(),
			),
		)
	}

	return html.Details(
		html.Class("search-result"),
		gomponents.If(open, html.Open()),
		html.Summary(
			html.A(
				html.Href(panelPath(c.Session)),
				extended.Get(panelPath(c.Session)),
				extended.Target(constant.PanelTarget),
				extended.Push("true"),
				gomponents.Text(name),
			),
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
