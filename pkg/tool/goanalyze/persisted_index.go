package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func persistedIndex(
	directory string,
	root string,
	patterns []string,
) (
	[]*packages.Package,
	map[string]bool,
	*face.Set,
	*unclosed_resource.Summaries,
) {
	reported, e := resolve.ListPackages(root, patterns...)
	errors.PanicOnError(e)
	result, _, e := resolve.LoadPackages(root, patterns...)
	errors.PanicOnError(e)
	loaded := resolve.PreferTestVariants(result)
	w := index.New(directory, root, face.Kind(), unclosed_resource.Kind())

	return loaded,
		reported,
		face.FromWorkspace(w, loaded),
		unclosed_resource.FromWorkspace(w, loaded)
}
