package layer

func (l *Layer) Holds(path string) bool {
	_, held := l.original[path]

	return held
}
