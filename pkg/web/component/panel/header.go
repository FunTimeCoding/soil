package panel

import "maragu.dev/gomponents"

func (p *Panel) Header(n ...gomponents.Node) *Panel {
	p.header = append(p.header, n...)

	return p
}
