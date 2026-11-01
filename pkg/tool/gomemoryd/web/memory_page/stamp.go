package memory_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func stamp(value string) gomponents.Node {
	return html.Small(
		gomponents.Text(
			library.FormatCompact(library.Parse(time.RFC3339, value)),
		),
	)
}
