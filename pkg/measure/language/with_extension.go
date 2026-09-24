package language

func (l *Language) WithExtension(v ...string) *Language {
	l.Extensions = append(l.Extensions, v...)

	return l
}
