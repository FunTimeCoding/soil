package example

import (
	"github.com/funtimecoding/soil/pkg/errors"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/preposition_name"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Census() {
	root, _ := lint.Root("")
	all, _, e := resolve.LoadPackages(root, goModConstant.AllPackages)
	errors.PanicOnError(e)
	loaded := resolve.PreferTestVariants(all)
	faces := face.NewWithIndex(
		index.New(index.DefaultDirectory(), root, face.Kind()),
		loaded,
	)
	results := output.NewResultsWithDirectory(root)

	for _, p := range loaded {
		preposition_name.Check(p, results, faces)
	}

	output.PrintResults(results.Entries, false)
}
