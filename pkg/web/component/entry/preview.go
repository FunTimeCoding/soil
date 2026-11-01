package entry

func (e *Entry) Preview(
	path string,
	target string,
) *Entry {
	e.previewPath = path
	e.previewTarget = target

	return e
}
