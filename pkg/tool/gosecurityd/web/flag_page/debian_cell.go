package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func DebianCell(
	r *status_result.Result,
	e error,
) gomponents.Node {
	if e != nil {
		return html.Small(gomponents.Text(e.Error()))
	}

	if len(r.Statuses) == 0 {
		return html.Small(gomponents.Text("no entries"))
	}

	items := make([]gomponents.Node, 0, len(r.Statuses))

	for _, v := range r.Statuses {
		text := fmt.Sprintf("%s %s: %s", v.Package, v.Release, v.Status)

		if v.NeverAffected {
			text = fmt.Sprintf("%s %s: never affected", v.Package, v.Release)
		} else if v.Fixed != "" {
			text = fmt.Sprintf(
				"%s %s: fixed in %s",
				v.Package,
				v.Release,
				v.Fixed,
			)
		}

		items = append(items, html.Li(gomponents.Text(text)))
	}

	return html.Ul(items...)
}
