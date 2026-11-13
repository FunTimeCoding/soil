package layout

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) container() gomponents.Node {
	if p.fullWindow {
		return html.Class("container-fluid")
	}

	return html.Class("container")
}
