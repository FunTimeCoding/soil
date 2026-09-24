package language

func (l *Language) WithBlockComment(
	open string,
	close string,
) *Language {
	l.BlockComments = append(l.BlockComments, &Block{Open: open, Close: close})

	return l
}
