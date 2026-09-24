package language

func (l *Language) WithSuffix(v ...string) *Language {
	l.Suffixes = append(l.Suffixes, v...)

	return l
}
