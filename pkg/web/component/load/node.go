package load

import (
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (l *Load) Node() gomponents.Node {
	return html.A(
		html.Href(l.path),
		gomponents.If(l.button, html.Role("button")),
		gomponents.If(l.class != "", html.Class(l.class)),
		gomponents.If(l.target != "", extended.Get(l.path)),
		gomponents.If(l.target != "", extended.Target(l.target)),
		gomponents.If(l.target != "", extended.Swap(l.swap)),
		gomponents.Text(l.label),
	)
}
