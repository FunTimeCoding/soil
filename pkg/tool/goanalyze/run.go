package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/option"
	"os"
)

func Run(o *option.Analyze) {
	results := Analyze(o)

	if output.PrintResults(results.Entries, o.Summary) {
		os.Exit(1)
	}
}
