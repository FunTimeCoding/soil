package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func versions(r *overview_result.Row) gomponents.Node {
	if r.Unread != nil {
		return gomponents.Group(nil)
	}

	items := make([]gomponents.Node, 0, len(r.Debian.Statuses))

	for _, s := range r.Debian.Statuses {
		current := make([]string, 0, len(s.Repositories))

		for _, archive := range slices.Sorted(maps.Keys(s.Repositories)) {
			current = append(
				current,
				fmt.Sprintf("%s %s", archive, s.Repositories[archive]),
			)
		}

		items = append(
			items,
			html.Li(
				gomponents.Text(
					fmt.Sprintf(
						"%s %s: %s; current %s",
						s.Package,
						s.Release,
						s.Status,
						join.CommaSpace(current),
					),
				),
			),
		)
	}

	return html.Details(
		html.Summary(gomponents.Text("Versions")),
		html.Ul(items...),
	)
}
