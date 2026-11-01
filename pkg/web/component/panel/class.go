package panel

func (p *Panel) Class(class string) *Panel {
	p.class = class

	return p
}
