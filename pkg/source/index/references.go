package index

import "github.com/funtimecoding/soil/pkg/source/index/xref"

func (w *Workspace) References() *xref.Index {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.references != nil {
		return w.references
	}

	w.ensure()
	directories := make(map[string]string, len(w.graph.Units))

	for path, u := range w.graph.Units {
		directories[path] = u.Directory
	}

	w.references = xref.New(w.refreshReferences(), directories)

	return w.references
}
