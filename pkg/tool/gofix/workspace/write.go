package workspace

func (w *Workspace) Write(
	path string,
	content []byte,
) error {
	w.layer.Write(path, content)

	return nil
}
