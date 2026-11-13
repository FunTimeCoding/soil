package plate

func (p *Plate) Name(text string) *Plate {
	p.name = text

	return p
}
