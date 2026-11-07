package layout

func (p *Page) HasCommandPalette() bool {
	return len(p.items)+len(p.commands) > 0
}
