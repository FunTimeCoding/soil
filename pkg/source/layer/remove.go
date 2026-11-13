package layer

func (l *Layer) Remove(path string) {
	l.touch(path)
	delete(l.written, path)
	l.removed[path] = true
	l.writes++
}
