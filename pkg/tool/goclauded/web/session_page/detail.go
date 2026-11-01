package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/component"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Detail(cruiseAlert string) gomponents.Node {
	d := p.view.Detail

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(p.Title())),
			component.LabelChips(p.view.Labels),
			html.Div(
				html.Class(constant.DetailHeadClass),
				p.identity(),
				gomponents.If(
					d.Description != "",
					html.P(gomponents.Text(d.Description)),
				),
			),
			p.tokens(),
			p.completions(),
			p.loads(),
			gomponents.If(
				d.Summary != "",
				gomponents.Group(
					[]gomponents.Node{
						html.H4(gomponents.Text("Summary")),
						html.Pre(gomponents.Text(d.Summary)),
					},
				),
			),
			html.H4(gomponents.Text("Pulses")),
			html.Div(
				extended.StreamSwap(constant.Pulse),
				Pulses(p.view.Pulses),
			),
			p.pulseForm(),
			html.H4(gomponents.Text("Cruise")),
			p.CruiseForm(cruiseAlert),
			html.H4(gomponents.Text("Session")),
			p.actions(),
		},
	)
}
