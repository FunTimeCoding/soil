package tile

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (t *Tile) Node() gomponents.Node {
	class := constant.TileClass

	if t.class != "" {
		class = join.Space(class, t.class)
	}

	return html.Div(
		html.Class(class),
		gomponents.If(
			t.terms != "",
			gomponents.Attr(constant.SearchAttribute, t.terms),
		),
		gomponents.Group(t.content),
	)
}
