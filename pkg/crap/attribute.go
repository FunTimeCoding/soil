package crap

import (
	"github.com/funtimecoding/soil/pkg/crap/attribution"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/coverage"
	"github.com/funtimecoding/soil/pkg/crap/index"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func Attribute(
	root string,
	i *index.Index,
	testPackages []string,
	coverPackages []string,
) *attribution.Matrix {
	result := attribution.New()
	work := system.TemporaryDirectory(constant.WorkPrefix)
	defer system.Remove(work)

	for _, p := range testPackages {
		directory := filepath.Join(work, join.Underscore(split.Slash(p)))
		system.MakeDirectory(directory)
		binary := coverage.Compile(root, directory, p, coverPackages)

		for _, t := range coverage.ListTests(binary) {
			profile := coverage.RunTestAlone(root, binary, t, directory)
			var keys []string

			for _, k := range coverage.CoveredKeys(coverage.Functions(root, profile)) {
				if i.ByKey(k) != nil {
					keys = append(keys, k)
				}
			}

			result.Add(join.Slash([]string{p, t}), keys)
		}
	}

	return result
}
