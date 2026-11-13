package plate

func (p *Plate) Headline(text string) *Plate {
	p.headline = text

	return p
}
