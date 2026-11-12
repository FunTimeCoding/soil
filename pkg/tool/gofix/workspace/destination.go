package workspace

func (w *Workspace) Destination(path string) string {
	if to, moved := w.moves[path]; moved {
		return to
	}

	return path
}
