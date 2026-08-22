package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func coverageToolRow(
	e *coverage.Server,
	t *coverage.Tool,
) gomponents.Node {
	state := "registered"

	if !t.Registered {
		state = "observed"

		if e.Registered > 0 {
			state = "retired"
		}
	}

	last := html.Td(gomponents.Text("-"))

	if !t.LastUsed.IsZero() {
		last = layout.TimeCell(t.LastUsed)
	}

	return html.Tr(
		html.Td(gomponents.Text(t.Name)),
		html.Td(gomponents.Text(state)),
		html.Td(gomponents.Text(fmt.Sprintf("%d", t.CallsRecent))),
		html.Td(gomponents.Text(fmt.Sprintf("%d", t.CallsTotal))),
		last,
	)
}
