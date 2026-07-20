package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func messageRow(m *message.Message) gomponents.Node {
	to := m.ToName

	if to == "" {
		to = "all"
	}

	return html.Tr(
		layout.TimeCell(m.CreatedAt),
		html.Td(gomponents.Text(m.FromName)),
		html.Td(gomponents.Text(to)),
		html.Td(gomponents.Text(m.Body)),
	)
}
