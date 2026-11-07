package policy

import "maps"

func (w *Walk) conditional(f func() bool) bool {
	base := w.base
	variables := maps.Clone(w.variables)
	result := f()

	if w.base != base {
		w.base = ""
	}

	for name, value := range w.variables {
		if before, okay := variables[name]; !okay || before != value {
			delete(w.variables, name)
		}
	}

	return result
}
