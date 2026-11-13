package workspace

func (w *Workspace) Read(path string) ([]byte, error) {
	return w.layer.Read(path)
}
