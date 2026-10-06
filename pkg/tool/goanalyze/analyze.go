package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"golang.org/x/tools/go/packages"
)

func Analyze(o *option.Analyze) *output.Results {
	patterns := o.Patterns

	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	var loaded []*packages.Package
	var reported map[string]bool
	var faces *face.Set
	var summaries *unclosed_resource.Summaries

	if o.Full || resolve.CoversMainModule(patterns) {
		loaded, reported, faces, summaries = fullIndex(o.Root, patterns)
	} else {
		directory := o.Index

		if directory == "" {
			directory = index.DefaultDirectory()
		}

		loaded, reported, faces, summaries = persistedIndex(
			directory,
			o.Root,
			patterns,
		)
	}

	result := output.NewResultsWithDirectory(o.Root)

	for _, p := range loaded {
		if reported[p.PkgPath] {
			check(p, result, o.Comment, faces, summaries)
		}
	}

	return result
}
