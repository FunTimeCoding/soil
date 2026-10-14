package index

func (w *Workspace) WithFingerprint(build string) *Workspace {
	w.build = build

	return w
}
