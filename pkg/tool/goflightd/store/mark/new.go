package mark

import "time"

func New(
	t time.Time,
	label string,
	note string,
) *Mark {
	return &Mark{Time: t, Label: label, Note: note}
}
