package example

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unread_receiver"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Census() {
	root, _ := lint.Root("")
	all, _, e := resolve.LoadPackages(root, "./...")
	errors.PanicOnError(e)
	loaded := resolve.PreferTestVariants(all)
	faces := face.NewWithIndex(
		index.New(index.DefaultDirectory(), root, face.Kind()),
		loaded,
	)
	results := output.NewResultsWithDirectory(root)

	for _, p := range loaded {
		unread_receiver.Check(p, results, faces)
	}

	output.PrintResults(results.Entries, false)
}
