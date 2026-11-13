package coverage_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/coverage"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func cells(e *coverage.Server) []gomponents.Node {
	name := e.Name

	if !e.Configured {
		name = fmt.Sprintf("%s (unconfigured)", name)
	}

	return []gomponents.Node{
		html.Td(gomponents.Text(name)),
		html.Td(gomponents.Text(ratio(e.UsedRecent, e.Registered))),
		html.Td(gomponents.Text(ratio(e.UsedTotal, e.Registered))),
		html.Td(gomponents.Textf("%d", e.CallsRecent)),
		html.Td(gomponents.Textf("%d", e.CallsTotal)),
		lastUsed(e.LastUsed),
	}
}
