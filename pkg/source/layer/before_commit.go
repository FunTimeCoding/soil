package layer

func (l *Layer) BeforeCommit(f func()) {
	l.beforeCommit = f
}
