package workspace

func (w *Workspace) Taken(path string) bool {
	return w.layer.Exists(path)
}
