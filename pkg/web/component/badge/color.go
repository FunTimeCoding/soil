package badge

func (b *Badge) Color(color string) *Badge {
	b.color = color

	return b
}
