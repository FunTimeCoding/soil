package byte_limit

func (l *Limit) Open(opening string) {
	l.opening = len(opening)
	l.empty = true
}
