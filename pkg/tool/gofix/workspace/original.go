package workspace

func (w *Workspace) Original(path string) []byte {
	return w.original[path]
}
