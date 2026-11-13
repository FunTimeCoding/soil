package layer

func (l *Layer) Writes() int {
	if l.parent == nil {
		return l.writes
	}

	return l.writes + l.parent.Writes()
}
