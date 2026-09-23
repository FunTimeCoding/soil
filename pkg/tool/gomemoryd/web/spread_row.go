package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func spreadRow(
	name string,
	s *token_summary.Spread,
) gomponents.Node {
	return html.Tr(
		html.Td(gomponents.Text(name)),
		html.Td(gomponents.Textf("%d", s.Median)),
		html.Td(gomponents.Textf("%d", s.Ninetieth)),
		html.Td(gomponents.Textf("%d", s.NinetyNinth)),
		html.Td(gomponents.Textf("%d", s.Maximum)),
		html.Td(gomponents.Textf("%d", s.Total)),
	)
}
