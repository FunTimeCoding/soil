package layout

func (p *Page) HasCommandPalette() bool {
	return len(p.items) > 0 || p.commands != nil
}
