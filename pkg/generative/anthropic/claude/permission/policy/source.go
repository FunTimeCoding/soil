package policy

import "slices"

func (w *Walk) source(
	allowed bool,
	name string,
) bool {
	if allowed && !slices.Contains(w.sources, name) {
		w.sources = append(w.sources, name)
	}

	return allowed
}
