package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"os"
)

func Lint(
	name string,
	root string,
	o *option.Lint,
) {
	r := Report(name, root, o)
	hasBlocked := output.Print(r, o.Summary)

	if o.Census {
		output.PrintCensus(r.Unchecked)
	}

	if hasBlocked {
		os.Exit(1)
	}
}
