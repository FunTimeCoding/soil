package language

func (l *Language) WithNested() *Language {
	l.Nested = true

	return l
}
