package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"os"
)

func Run(o *option.Analyze) {
	patterns := o.Patterns

	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	loaded := load(patterns)
	results := output.NewResults()

	for _, p := range loaded {
		check(p, results)
	}

	hasBlocked := output.PrintResults(results.Entries, o.Summary)

	if hasBlocked {
		os.Exit(1)
	}
}
