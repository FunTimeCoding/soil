package example

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/test_article"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Census() {
	root, _ := lint.Root("")
	all, _, e := resolve.LoadPackages(root, "./...")
	errors.PanicOnError(e)
	results := output.NewResultsWithDirectory(root)

	for _, p := range resolve.PreferTestVariants(all) {
		test_article.Check(p, results)
	}

	output.PrintResults(results.Entries, false)
}
