package layer

func (l *Layer) Write(
	path string,
	content []byte,
) {
	l.touch(path)
	delete(l.removed, path)
	l.written[path] = content
	l.writes++
}
