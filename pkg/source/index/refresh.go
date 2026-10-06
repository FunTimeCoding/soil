package index

import "github.com/funtimecoding/soil/pkg/source/constant"

func (w *Workspace) refresh() {
	g := w.graph
	w.fingerprints = make(map[string]string, len(g.Nodes))
	w.facts = make(map[string]map[string]any, len(w.kinds))

	for _, k := range w.kinds {
		w.facts[k.Name] = make(map[string]any, len(g.Nodes))
	}

	done := make(map[string]bool, len(g.Nodes))
	failed := make(map[string]bool)
	files := make(map[string]string, len(g.Nodes))

	for path, n := range g.Nodes {
		files[path] = w.fileHash(n.Files)
	}

	for {
		var misses []string

		for _, path := range g.Order {
			if done[path] || failed[path] || !ready(g, path, w.fingerprints) {
				continue
			}

			if w.read(path, key(g, g.Nodes[path], files[path], w.fingerprints)) {
				done[path] = true

				continue
			}

			misses = append(misses, path)
		}

		if len(misses) == 0 {
			return
		}

		if len(misses) > constant.IndexBulkThreshold {
			misses = remaining(g, done, failed)
		}

		w.compute(misses, files, done, failed)
	}
}
