package workspace

func (w *Workspace) BeforeCommit(f func()) {
	w.beforeCommit = f
}
