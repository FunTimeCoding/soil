package layout

func (p *Page) WithServerSide(source string) *Page {
	p.serverSide = source

	return p
}
