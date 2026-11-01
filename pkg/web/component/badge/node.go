package badge

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (b *Badge) Node() gomponents.Node {
	class := constant.BadgeClass

	if b.quiet {
		class = join.Space(class, constant.BadgeQuietClass)
	}

	return html.Span(
		html.Class(class),
		gomponents.If(
			b.color != "",
			html.Style(join.Empty(constant.BadgeColorVariable, ": ", b.color)),
		),
		gomponents.Text(b.label),
	)
}
