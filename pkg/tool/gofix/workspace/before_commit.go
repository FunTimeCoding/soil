package workspace

func (w *Workspace) BeforeCommit(f func()) {
	w.layer.BeforeCommit(f)
}
