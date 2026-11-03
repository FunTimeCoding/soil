package memory_page

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func stamp(value string) gomponents.Node {
	return html.Small(
		gomponents.Text(
			soilTime.FormatCompact(soilTime.Parse(time.RFC3339, value)),
		),
	)
}
