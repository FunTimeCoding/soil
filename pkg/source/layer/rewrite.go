package layer

func (l *Layer) Rewrite(
	path string,
	expected []byte,
	content []byte,
) {
	l.Guard(path, expected)
	l.Write(path, content)
}
