package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageDetail(e *coverage.Server) gomponents.Node {
	return html.Details(
		html.Summary(gomponents.Text(e.Name)),
		coverageToolTable(e),
	)
}
