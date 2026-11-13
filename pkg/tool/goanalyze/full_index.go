package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
	"io"
	"slices"
)

func fullIndex(
	root string,
	raw []*packages.Package,
	handle *index.Workspace,
	notice io.Writer,
) ([]*packages.Package, *face.Set, *unclosed_resource.Summaries) {
	loaded := resolve.PreferTestVariants(raw)
	reached := loadReached(root, loaded, notice)

	return loaded,
		face.NewWithIndex(handle, loaded),
		unclosed_resource.NewSummaries(
			slices.Concat(loaded, reached),
			index.Externals[*fact.Summary](handle, unclosed_resource.Kind())...,
		)
}
