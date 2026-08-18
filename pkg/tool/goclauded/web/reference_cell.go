package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func referenceCell(
	entry *context_load.Load,
	base string,
) gomponents.Node {
	if entry.Reference == "" ||
		entry.Kind == constant.LoadKindMode ||
		base == "" {
		return html.Td(html.Small(gomponents.Text(entry.Reference)))
	}

	return html.Td(
		html.Small(
			html.A(
				gomponents.Attr("href", join.Empty(base, entry.Reference)),
				gomponents.Text(entry.Reference),
			),
		),
	)
}
