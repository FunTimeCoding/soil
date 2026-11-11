package budget

func (b *Budget) Charge(item string) {
	b.spent += b.cost(item)
	b.empty = false
	b.rendered = true
}
