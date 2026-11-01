package badge

func (b *Badge) Quiet() *Badge {
	b.quiet = true

	return b
}
