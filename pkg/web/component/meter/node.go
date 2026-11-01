package meter

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (m *Meter) Node() gomponents.Node {
	if len(m.label) == 0 && m.shown == "" && m.icon == nil &&
		m.title == "" && m.class == "" {
		return m.bar()
	}

	class := constant.MeterClass

	if m.class != "" {
		class = join.Space(class, m.class)
	}

	return html.Div(
		html.Class(class),
		gomponents.If(
			m.title != "",
			gomponents.Attr(constant.TooltipAttribute, m.title),
		),
		gomponents.If(
			m.icon != nil,
			html.Span(html.Class(constant.MeterIconClass), m.icon),
		),
		gomponents.If(
			len(m.label) > 0,
			html.Span(
				html.Class(constant.MeterLabelClass),
				gomponents.Group(m.label),
			),
		),
		m.bar(),
		gomponents.If(
			m.shown != "",
			html.Span(
				html.Class(constant.MeterValueClass),
				gomponents.Text(m.shown),
			),
		),
	)
}
