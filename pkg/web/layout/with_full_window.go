package layout

func (p *Page) WithFullWindow() *Page {
	p.fullWindow = true

	return p
}
