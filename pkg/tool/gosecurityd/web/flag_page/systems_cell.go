package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func systemsCell(c *affected_count.Count) gomponents.Node {
	if c == nil || c.Systems == 0 {
		return html.Td(html.Small(gomponents.Text("none")))
	}

	return html.Td(
		gomponents.Text(fmt.Sprint(c.Systems)),
		html.Br(),
		html.Small(gomponents.Text(fmt.Sprintf("%d upgrade fixes", c.Upgrade))),
	)
}
