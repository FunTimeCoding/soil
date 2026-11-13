package memory_page

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func cells(
	name string,
	description string,
	tags []string,
) []gomponents.Node {
	return []gomponents.Node{
		html.Td(gomponents.Text(name)),
		layout.ClampCell(description),
		html.Td(TagChips(tags)),
	}
}
