package entry

func (e *Entry) Swap(
	target string,
	style string,
) *Entry {
	e.target = target
	e.swap = style

	return e
}
