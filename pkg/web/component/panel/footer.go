package panel

import "maragu.dev/gomponents"

func (p *Panel) Footer(n ...gomponents.Node) *Panel {
	p.footer = append(p.footer, n...)

	return p
}
