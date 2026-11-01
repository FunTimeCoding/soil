package floor_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Region(
	f floor.Floor,
	links map[guest.Guest]string,
) gomponents.Node {
	if len(f.Nodes) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No hypervisors seen yet - first poll pending."),
		)
	}

	var sections []gomponents.Node

	for _, n := range f.Nodes {
		sections = append(sections, section(n, f.Guests, links)...)
	}

	return html.Div(sections...)
}
