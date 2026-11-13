package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
)

func Analyze(o *option.Analyze) *output.Results {
	patterns := analyzedPatterns(o)
	loadedPatterns := patterns

	if whole(o, patterns) {
		loadedPatterns = resolve.WithMainModule(patterns)
	}

	directory := o.Index

	if directory == "" {
		directory = index.DefaultDirectory()
	}

	raw, reported := load(o.Root, patterns, loadedPatterns)

	return AnalyzeLoaded(
		o,
		raw,
		reported,
		index.New(directory, o.Root, Kinds()...),
	)
}
