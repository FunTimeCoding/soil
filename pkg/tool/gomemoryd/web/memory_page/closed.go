package memory_page

import "maragu.dev/gomponents"

func (p *Page) Closed() gomponents.Node {
	return p.row().Response()
}
