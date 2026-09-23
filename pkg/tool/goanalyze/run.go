package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"os"
)

func Run(o *option.Analyze) {
	patterns := o.Patterns

	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	loaded := load(o.Root, patterns)
	results := output.NewResultsWithDirectory(o.Root)
	faces := face.New(loaded)
	summaries := unclosed_resource.NewSummaries(loaded)

	for _, p := range loaded {
		check(p, results, o.Comment, faces, summaries)
	}

	hasBlocked := output.PrintResults(results.Entries, o.Summary)

	if hasBlocked {
		os.Exit(1)
	}
}
