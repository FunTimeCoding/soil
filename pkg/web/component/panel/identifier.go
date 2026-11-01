package panel

func (p *Panel) Identifier(identifier string) *Panel {
	p.identifier = identifier

	return p
}
