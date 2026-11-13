package concurrent_write

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"path/filepath"
	"slices"
)

func Reached(
	graphs []*module_graph.Graph,
	packagePaths []string,
	written []string,
	changed []string,
) []string {
	if len(packagePaths) == 0 || len(changed) == 0 {
		return changed
	}

	kept := map[string]bool{}

	for _, p := range written {
		kept[filepath.Dir(p)] = true
	}

	for _, g := range graphs {
		for _, u := range module_graph.ReachUnits(g, packagePaths) {
			kept[u.Directory] = true
		}
	}

	var result []string
	files := []string{goModConstant.ModFile, goModConstant.SumFile}

	for _, p := range changed {
		if kept[filepath.Dir(p)] || slices.Contains(files, filepath.Base(p)) {
			result = append(result, p)
		}
	}

	return result
}
