package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"golang.org/x/tools/go/packages"
	"slices"
)

func fullIndex(
	root string,
	patterns []string,
) (
	[]*packages.Package,
	map[string]bool,
	*face.Set,
	*unclosed_resource.Summaries,
) {
	loaded, reported, reached := load(root, patterns)

	return loaded,
		reported,
		face.New(loaded),
		unclosed_resource.NewSummaries(slices.Concat(loaded, reached))
}
