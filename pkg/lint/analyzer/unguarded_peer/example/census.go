package example

import (
	"github.com/funtimecoding/soil/pkg/errors"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unguarded_peer"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Census() {
	root, _ := lint.Root("")
	all, _, e := resolve.LoadPackages(root, goModConstant.AllPackages)
	errors.PanicOnError(e)
	results := output.NewResultsWithDirectory(root)

	for _, p := range resolve.PreferTestVariants(all) {
		unguarded_peer.Check(p, results)
	}

	output.PrintResults(results.Entries, false)
}
