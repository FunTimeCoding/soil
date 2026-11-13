package layer

func (l *Layer) Guard(
	path string,
	expected []byte,
) {
	if _, guarded := l.expected[path]; !guarded {
		l.expected[path] = expected
	}
}
