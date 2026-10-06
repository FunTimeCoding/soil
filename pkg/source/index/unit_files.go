package index

import "path/filepath"

func (w *Workspace) UnitFiles() map[string]bool {
	result := make(map[string]bool)

	for _, u := range w.graph.Units {
		for _, f := range u.Files {
			if full, e := filepath.Abs(f); e == nil {
				result[full] = true
			}
		}
	}

	return result
}
