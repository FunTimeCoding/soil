package go_mod

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/go_mod/dependency"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"sort"
	"strconv"
)

func ListDependencies(
	path string,
	verbose bool,
) []*dependency.Dependency {
	r := run.New()
	r.Directory = path
	r.Verbose = verbose
	r.Start(
		library.Go,
		library.List,
		constant.DependenciesArgument,
		constant.FormatArgument,
		constant.ModuleTemplate,
		constant.AllPackages,
	)
	seen := map[string]bool{}
	var result []*dependency.Dependency

	for _, line := range split.NewLine(r.OutputString) {
		fields := split.Space(line)

		if len(fields) != 4 ||
			fields[0] == strconv.FormatBool(true) ||
			seen[fields[1]] {
			continue
		}

		seen[fields[1]] = true
		result = append(result, dependency.New(fields[1], fields[2], fields[3]))
	}

	sort.Slice(
		result,
		func(i int, j int) bool {
			return result[i].Path < result[j].Path
		},
	)

	return result
}
