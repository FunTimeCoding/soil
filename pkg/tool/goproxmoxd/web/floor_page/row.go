package floor_page

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor/guest"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func row(
	g guest.Guest,
	link string,
) gomponents.Node {
	running := g.Status == proxmoxConstant.RunningStatus
	color := webConstant.QuietColor

	if running {
		color = webConstant.OkayColor
	}

	return html.Tr(
		gomponents.If(!running, html.Class(constant.StoppedClass)),
		html.Td(status_dot.New(color).Title(g.Status).Node()),
		html.Td(gomponents.Textf("%d", g.Identifier)),
		html.Td(
			html.A(
				html.Href(link),
				html.Target("_blank"),
				gomponents.Text(g.Name),
			),
		),
		html.Td(html.Class(constant.LoadClass), gomponents.Text(load(g))),
		html.Td(
			gomponents.If(
				g.Unbacked,
				badge.New("unbacked").Color(webConstant.BadColor).Node(),
			),
		),
	)
}
