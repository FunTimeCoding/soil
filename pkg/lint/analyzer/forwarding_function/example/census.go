package example

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/forwarding_function"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Census() {
	all, _, e := resolve.LoadPackages(soil.CurrentDirectory, "./...")
	errors.PanicOnError(e)
	results := output.NewResultsWithDirectory(soil.CurrentDirectory)

	for _, p := range resolve.PreferTestVariants(all) {
		forwarding_function.Check(p, results)
	}

	output.PrintResults(results.Entries, false)
}
