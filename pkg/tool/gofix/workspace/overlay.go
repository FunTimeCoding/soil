package workspace

func (w *Workspace) Overlay() map[string][]byte {
	return w.layer.Overlay()
}
