package load

func (l *Load) Swap(
	target string,
	style string,
) *Load {
	l.target = target
	l.swap = style

	return l
}
