package commit

import (
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"path/filepath"
	"slices"
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

	kept := map[string]bool{}

	for _, out := range outs {
		for _, p := range out.Paths() {
			kept[filepath.Dir(p)] = true
		}

		for _, o := range out.Renames() {
			if o.Target != "" {
				kept[filepath.Dir(o.Target)] = true
			}
		}
	}

	for _, root := range roots {
		g := s.workspace.Workspace(root).Graph()

		for _, u := range load.ReachUnits(g, packagePaths) {
			kept[u.Directory] = true
		}
	}

	var result []string
	files := []string{go_modConstant.ModFile, go_modConstant.SumFile}

	for _, p := range changed {
		name := filepath.Base(p)

		if kept[filepath.Dir(p)] || slices.Contains(files, name) {
			result = append(result, p)
		}
	}

	return result
}
