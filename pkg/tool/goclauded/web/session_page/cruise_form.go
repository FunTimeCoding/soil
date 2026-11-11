package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) CruiseForm(alert string) gomponents.Node {
	setting := p.view.Cruise

	return gomponents.Group(
		[]gomponents.Node{
			cruiseEntry(p.view.Detail.Identifier, setting, p.view.CruiseGrace).
				Identifier(constant.CruiseAction).
				Alert(alert).
				Node(),
			html.P(
				html.Class(webConstant.AbsentClass),
				gomponents.Text(cruiseState(setting, p.view.CruiseGrace)),
			),
		},
	)
}
