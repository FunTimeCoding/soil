package index

func (w *Workspace) Mark(
	name string,
	key string,
) {
	w.store.Write(w.marksStore(name), key, true)
}
