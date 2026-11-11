package budget

func (b *Budget) Fit(item string) bool {
	if b.spent+b.cost(item) > b.limit {
		return false
	}

	b.Charge(item)

	return true
}
