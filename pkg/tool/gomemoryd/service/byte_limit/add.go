package byte_limit

func (l *Limit) Add(item string) {
	l.used += l.cost(item)
	l.empty = false
	l.rendered = true
}
