package pager

func (p *Pager) Position(text string) *Pager {
	p.position = text

	return p
}
