package commit

import (
	"github.com/funtimecoding/soil/pkg/source/concurrent_write"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) reachedChanges(
	roots []string,
	packagePaths []string,
	outs []*sink.Sink,
	changed []string,
) []string {
	if len(packagePaths) == 0 || len(changed) == 0 {
		return changed
	}

	var written []string

	for _, out := range outs {
		written = append(written, out.Paths()...)

		for _, o := range out.Renames() {
			if o.Target != "" {
				written = append(written, o.Target)
			}
		}
	}

	var graphs []*module_graph.Graph

	for _, root := range roots {
		graphs = append(graphs, s.workspace.Workspace(root).Graph())
	}

	return concurrent_write.Reached(graphs, packagePaths, written, changed)
}
