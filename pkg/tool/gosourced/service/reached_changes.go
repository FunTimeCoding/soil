package service

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
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
		for _, u := range reachUnits(s.workspace(root).Graph(), packagePaths) {
			kept[u.Directory] = true
		}
	}

	var result []string

	for _, p := range changed {
		name := filepath.Base(p)

		if kept[filepath.Dir(p)] ||
			slices.Contains([]string{constant.ModFile, constant.SumFile}, name) {
			result = append(result, p)
		}
	}

	return result
}
