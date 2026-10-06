package workspace

func (w *Workspace) Overlay() map[string][]byte {
	if len(w.overlay) == 0 {
		return nil
	}

	return w.overlay
}
