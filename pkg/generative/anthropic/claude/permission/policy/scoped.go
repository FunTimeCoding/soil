package policy

import "maps"

func (w *Walk) scoped(f func() bool) bool {
	base := w.base
	variables := maps.Clone(w.variables)
	result := f()
	w.base = base
	w.variables = variables

	return result
}
