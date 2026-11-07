package policy

import "slices"

func (w *Walk) block(label string) bool {
	if !slices.Contains(w.blockers, label) {
		w.blockers = append(w.blockers, label)
	}

	return false
}
