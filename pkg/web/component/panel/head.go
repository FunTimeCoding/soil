package panel

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Panel) head() gomponents.Node {
	if p.title == "" {
		if len(p.header) == 0 {
			return nil
		}

		return html.Header(p.header...)
	}

	return html.Header(
		gomponents.If(len(p.header) > 0, html.Class(constant.HeadingRowClass)),
		html.Strong(gomponents.Text(p.title)),
		gomponents.Group(p.header),
	)
}
