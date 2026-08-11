package layout

func (p *Page) Clone() *Page {
	return &Page{
		identity:        p.identity,
		theme:           p.theme,
		style:           p.style,
		items:           p.items,
		navigation:      p.navigation,
		script:          p.script,
		footer:          p.footer,
		liveEndpoint:    p.liveEndpoint,
		paletteEndpoint: p.paletteEndpoint,
	}
}
