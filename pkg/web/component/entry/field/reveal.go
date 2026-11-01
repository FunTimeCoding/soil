package field

func (f *Field) Reveal(
	path string,
	target string,
) *Field {
	f.revealPath = path
	f.revealTarget = target

	return f
}
