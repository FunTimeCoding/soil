package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func applyNaming(
	fileSet *token.FileSet,
	all []*packages.Package,
	violations []Violation,
	o *option.Fix,
	w *workspace.Workspace,
	r *output.Results,
	covered map[string]bool,
) {
	reached, okay := reachReplacing(o, violations, r)

	if !okay {
		return
	}

	edits := BuildAllEdits(fileSet, all, violations, r)
	editThrough(fileSet, edits, o.Root, w)
	loaded := BuildLoadedFiles(all)

	for f := range covered {
		loaded[f] = true
	}

	referencesThrough(violations, loaded, o.Root, r, w)

	for _, m := range reached {
		editThrough(m.fileSet, m.edits, m.directory, w)

		for _, c := range m.concerns {
			r.AddConcern(c)
		}
	}
}
