package workspace

import "os"

func (w *Workspace) Write(
	path string,
	content []byte,
) error {
	if _, seen := w.original[path]; !seen {
		before, e := os.ReadFile(path)

		if e != nil {
			return e
		}

		w.original[path] = before
		w.order = append(w.order, path)
	}

	w.overlay[path] = content
	w.writes++

	return nil
}
