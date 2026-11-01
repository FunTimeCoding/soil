package panel

import "maragu.dev/gomponents"

func (p *Panel) Content(n ...gomponents.Node) *Panel {
	p.content = append(p.content, n...)

	return p
}
