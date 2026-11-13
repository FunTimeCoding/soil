package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"golang.org/x/tools/go/packages"
)

func report(
	o *option.Analyze,
	loaded []*packages.Package,
	reported map[string]bool,
	faces *face.Set,
	summaries *unclosed_resource.Summaries,
) *output.Results {
	result := output.NewResultsWithDirectory(o.Root)

	for _, p := range loaded {
		if reported[p.PkgPath] {
			check(p, result, o.Comment, faces, summaries)
		}
	}

	return result
}
