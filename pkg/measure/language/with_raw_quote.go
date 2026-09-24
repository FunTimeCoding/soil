package language

func (l *Language) WithRawQuote(v ...string) *Language {
	l.RawQuotes = append(l.RawQuotes, v...)

	return l
}
