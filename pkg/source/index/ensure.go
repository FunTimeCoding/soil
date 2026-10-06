package index

func (w *Workspace) ensure() {
	if w.fingerprints == nil {
		w.refresh()
	}
}
