package statistic_page

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
		numericCell(s.Median),
		numericCell(s.Ninetieth),
		numericCell(s.NinetyNinth),
		numericCell(s.Maximum),
		numericCell(s.Total),
	)
}
