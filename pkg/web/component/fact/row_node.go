package fact

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) RowNode(
	label string,
	value gomponents.Node,
) *Fact {
	f.rows = append(
		f.rows,
		html.Tr(
			html.Th(html.Scope(constant.RowScope), gomponents.Text(label)),
			html.Td(value),
		),
	)

	return f
}
