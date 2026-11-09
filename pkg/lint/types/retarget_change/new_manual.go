package retarget_change

func NewManual(
	line int,
	before string,
	reason string,
) *Change {
	return &Change{Line: line, Old: before, Reason: reason}
}
