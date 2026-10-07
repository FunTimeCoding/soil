package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"github.com/funtimecoding/soil/pkg/source/index/record"
)

func (w *Workspace) walkExternals(k *kind.Kind) ([]any, []string) {
	var queue []string

	for _, n := range w.graph.Nodes {
		queue = append(queue, n.Imports...)
	}

	seen := make(map[string]bool)
	var result []any
	var missing []string

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		if seen[path] {
			continue
		}

		seen[path] = true

		if _, workspace := w.graph.Nodes[path]; workspace {
			continue
		}

		key := externalKey(path, externalVersion(path, w.graph.Requirements))
		r, found := w.fetch(
			constant.IndexImportsKind,
			path,
			key,
			record.NewExternal(nil),
		)
		v, held := w.fetch(
			kindStore(constant.IndexExternalKind, k),
			path,
			kindKey(key, k),
			k.Value(),
		)

		if !found || !held {
			missing = append(missing, path)

			continue
		}

		result = append(result, v)
		queue = append(queue, r.(*record.External).Imports...)
	}

	return result, missing
}
