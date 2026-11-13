package index

func (w *Workspace) Marked(
	name string,
	key string,
) bool {
	return w.store.Has(w.marksStore(name), key)
}
