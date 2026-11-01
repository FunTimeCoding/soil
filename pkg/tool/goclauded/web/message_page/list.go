package message_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/web/component/pager"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(
	messages []message.Message,
	page int,
	more bool,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(messages))

	for _, m := range messages {
		to := m.ToName

		if to == "" {
			to = "all"
		}

		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(m.CreatedAt),
				html.Td(gomponents.Text(m.FromName)),
				html.Td(gomponents.Text(to)),
				html.Td(gomponents.Text(m.Body)),
			),
		)
	}

	pages := pager.New()

	if page > 1 {
		pages.Previous(Path(page-1), "← Newer")
	}

	if more {
		pages.Next(Path(page+1), "Older →")
	}

	return gomponents.Group(
		[]gomponents.Node{
			table.New("Time", "From", "To", "Message").
				Rows(rows).
				Empty("No messages.").
				Node(),
			pages.Node(),
		},
	)
}
