package retarget_change

func New(
	line int,
	before string,
	after string,
) *Change {
	return &Change{Line: line, Before: before, After: after}
}
