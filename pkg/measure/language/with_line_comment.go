package language

func (l *Language) WithLineComment(v ...string) *Language {
	l.LineComments = append(l.LineComments, v...)

	return l
}
