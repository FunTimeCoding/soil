package collector

import "slices"

func withLabel(
	label []string,
	extra ...string,
) []string {
	return slices.Concat(label, extra)
}
