package budget

func (b *Budget) Open(opening string) {
	b.opening = len(opening)
	b.empty = true
}
