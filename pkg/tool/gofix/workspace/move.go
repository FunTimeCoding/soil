package workspace

func (w *Workspace) Move(
	from string,
	to string,
) {
	w.layer.Move(from, to)
}
