package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func row(e *coverage.Server) gomponents.Node {
	if len(e.Tools) == 0 {
		return html.Tr(gomponents.Group(cells(e)))
	}

	return opening(e).Node()
}
