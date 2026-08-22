package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageRow(e *coverage.Server) gomponents.Node {
	name := e.Name

	if !e.Configured {
		name = fmt.Sprintf("%s (unconfigured)", name)
	}

	return html.Tr(
		html.Td(gomponents.Text(name)),
		html.Td(gomponents.Text(coverageRatio(e.UsedRecent, e.Registered))),
		html.Td(gomponents.Text(coverageRatio(e.UsedTotal, e.Registered))),
		html.Td(gomponents.Text(fmt.Sprintf("%d", e.CallsRecent))),
		html.Td(gomponents.Text(fmt.Sprintf("%d", e.CallsTotal))),
		coverageTimeCell(e),
	)
}
