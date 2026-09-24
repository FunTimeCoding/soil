package measure

import (
	"github.com/funtimecoding/soil/pkg/console"
	"sort"
)

func printPaths(
	header string,
	paths []string,
) {
	if len(paths) == 0 {
		return
	}

	console.Line()
	console.Line(header)
	sorted := make([]string, len(paths))
	copy(sorted, paths)
	sort.Strings(sorted)

	for _, p := range sorted {
		console.Line(p)
	}
}
