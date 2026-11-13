package layer

func (l *Layer) Exists(path string) bool {
	_, e := l.Read(path)

	return e == nil
}
