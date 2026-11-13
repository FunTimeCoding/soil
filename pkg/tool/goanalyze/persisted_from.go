package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func persistedFrom(
	raw []*packages.Package,
	handle *index.Workspace,
) ([]*packages.Package, *face.Set, *unclosed_resource.Summaries) {
	loaded := resolve.PreferTestVariants(raw)

	return loaded,
		face.FromWorkspace(handle, loaded),
		unclosed_resource.FromWorkspace(handle, loaded)
}
