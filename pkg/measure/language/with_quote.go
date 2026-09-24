package language

func (l *Language) WithQuote(v ...string) *Language {
	l.Quotes = append(l.Quotes, v...)

	return l
}
