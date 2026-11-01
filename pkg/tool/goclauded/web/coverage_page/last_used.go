package coverage_page

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func lastUsed(t time.Time) gomponents.Node {
	if t.IsZero() {
		return html.Td(gomponents.Text("-"))
	}

	return layout.TimeCell(t)
}
