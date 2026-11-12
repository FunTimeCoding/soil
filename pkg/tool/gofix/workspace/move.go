package workspace

func (w *Workspace) Move(
	from string,
	to string,
) {
	w.moves[from] = to
}
