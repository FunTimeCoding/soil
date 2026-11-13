package plate

func (p *Plate) Caption(text string) *Plate {
	p.caption = text

	return p
}
