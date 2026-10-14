package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"golang.org/x/tools/go/packages"
	"slices"
)

func fullIndex(
	directory string,
	root string,
	patterns []string,
) (
	[]*packages.Package,
	map[string]bool,
	*face.Set,
	*unclosed_resource.Summaries,
) {
	loaded, reported, reached := load(root, patterns)
	w := index.New(directory, root, unclosed_resource.Kind())

	return loaded,
		reported,
		face.New(loaded),
		unclosed_resource.NewSummaries(
			slices.Concat(loaded, reached),
			index.Externals[*fact.Summary](w, unclosed_resource.Kind())...,
		)
}
