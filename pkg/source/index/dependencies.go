package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/record"
)

func (w *Workspace) Dependencies(path string) (map[string]bool, bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	result := map[string]bool{}
	complete := true
	queue := []string{path}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if n, workspace := w.graph.Nodes[p]; workspace {
			for _, i := range n.Imports {
				if !result[i] {
					result[i] = true
					queue = append(queue, i)
				}
			}

			continue
		}

		r, found := w.fetch(
			constant.IndexImportsKind,
			p,
			externalKey(p, externalVersion(p, w.graph.Requirements)),
			record.NewExternal(nil),
		)

		if !found {
			if p != path {
				complete = false
			}

			continue
		}

		for _, i := range r.(*record.External).Imports {
			if !result[i] {
				result[i] = true
				queue = append(queue, i)
			}
		}
	}

	return result, complete
}
