package language

func (l *Language) WithShebang(v ...string) *Language {
	l.Shebangs = append(l.Shebangs, v...)

	return l
}
