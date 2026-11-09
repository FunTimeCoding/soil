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
			cruiseEntry(p.view.Detail.Identifier, setting).
				Identifier(constant.CruiseAction).
				Alert(alert).
				Node(),
			gomponents.If(
				setting.PromotedAt != nil,
				html.P(
					html.Class(webConstant.AbsentClass),
					gomponents.Textf(
						"woken %d times, last %s",
						setting.Woken,
						promoted(setting.PromotedAt),
					),
				),
			),
		},
	)
}
