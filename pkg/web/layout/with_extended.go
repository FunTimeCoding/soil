package layout

func (p *Page) WithExtended(source string) *Page {
	p.extended = source

	return p
}
