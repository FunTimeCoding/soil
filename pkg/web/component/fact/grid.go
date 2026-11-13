package fact

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) grid(items []gomponents.Node) gomponents.Node {
	if f.columns == 0 {
		return html.Div(
			html.Class(constant.GlanceClass),
			gomponents.Group(items),
		)
	}

	return html.Div(
		html.Class(
			join.Space(constant.GlanceClass, constant.GlanceColumnsClass),
		),
		html.Style(
			fmt.Sprintf("%s: %d", constant.GlanceColumnsVariable, f.columns),
		),
		gomponents.Group(items),
	)
}
