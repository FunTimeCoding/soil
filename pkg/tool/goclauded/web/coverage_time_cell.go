package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageTimeCell(e *coverage.Server) gomponents.Node {
	if e.LastUsed.IsZero() {
		return html.Td(gomponents.Text("-"))
	}

	return layout.TimeCell(e.LastUsed)
}
