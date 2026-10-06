package workspace

func (w *Workspace) Current(path string) []byte {
	return w.overlay[path]
}
