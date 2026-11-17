package byte_limit

func (l *Limit) Fit(item string) bool {
	if l.used+l.cost(item) > l.limit {
		return false
	}

	l.Add(item)

	return true
}
