package layer

import (
	"github.com/funtimecoding/soil/pkg/source/concurrent_write"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"slices"
)

func (l *Layer) Commit(graphs []*module_graph.Graph) ([]string, error) {
	touched := l.touched()

	if len(touched) == 0 {
		return nil, nil
	}

	if l.beforeCommit != nil {
		l.beforeCommit()
	}

	var roots []string

	for _, root := range l.roots {
		if slices.ContainsFunc(
			touched,
			func(path string) bool {
				return system.InsideDirectory(root, path)
			},
		) {
			roots = append(roots, root)
		}
	}

	changed := concurrent_write.Reached(
		graphs,
		packagePaths(graphs, touched),
		touched,
		l.before.Changed(roots...),
	)

	if refused := slices.Concat(changed, l.stale()); len(refused) > 0 {
		return refused, nil
	}

	for _, path := range touched {
		content, e := l.Read(path)

		if e != nil {
			continue
		}

		system.EnsurePathExists(filepath.Dir(path))

		if f := os.WriteFile(path, content, 0644); f != nil {
			return nil, f
		}
	}

	for _, path := range touched {
		if _, e := l.Read(path); e == nil || !system.FileExists(path) {
			continue
		}

		if f := os.Remove(path); f != nil {
			return nil, f
		}
	}

	return nil, nil
}
