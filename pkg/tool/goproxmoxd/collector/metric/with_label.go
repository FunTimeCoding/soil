package metric

import "slices"

func WithLabel(
	label []string,
	extra ...string,
) []string {
	return slices.Concat(label, extra)
}
