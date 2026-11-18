package target

func (t *Target) Close() {
	t.Guard.Close()
}
