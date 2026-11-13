package fact

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) group(
	title string,
	items []gomponents.Node,
) gomponents.Node {
	return html.Div(
		html.Class(constant.GlanceGroupClass),
		gomponents.If(title != "", html.H5(gomponents.Text(title))),
		f.grid(items),
	)
}
