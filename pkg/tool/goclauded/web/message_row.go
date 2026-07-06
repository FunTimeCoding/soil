package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func messageRow(m *message.Message) gomponents.Node {
	to := m.ToName

	if to == "" {
		to = "all"
	}

	return html.Tr(
		html.Td(html.Small(gomponents.Text(m.CreatedAt.Format("15:04")))),
		html.Td(gomponents.Text(m.FromName)),
		html.Td(gomponents.Text(to)),
		html.Td(gomponents.Text(m.Body)),
	)
}
