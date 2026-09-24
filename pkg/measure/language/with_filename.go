package language

func (l *Language) WithFilename(v ...string) *Language {
	l.Filenames = append(l.Filenames, v...)

	return l
}
