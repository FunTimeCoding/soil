package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"slices"
)

func targetedFix(
	o *option.Fix,
	patterns []string,
	reported map[string]bool,
	faces *face.Set,
	violations []violation,
	handle *index.Workspace,
	w *workspace.Workspace,
	r *output.Results,
) {
	i := handle.References()

	if len(i.Unindexed()) > 0 {
		fixModule(o, patterns, reported, faces, w, r)

		return
	}

	extended := slices.Clone(patterns)

	for _, v := range violations {
		t, okay := xref.Target(v.object)

		if !okay {
			fixModule(o, patterns, reported, faces, w, r)

			return
		}

		for _, unit := range i.Referencing(t) {
			extended = append(extended, i.Directory(unit))
		}
	}

	extended = slices.Compact(slices.Sorted(slices.Values(extended)))

	if 2*len(extended) > i.Count() {
		fixModule(o, patterns, reported, faces, w, r)

		return
	}

	all, fileSet := loadThrough(o.Root, extended, w)
	again := FindViolations(all, reported, faces)

	if len(again) > 0 {
		applyNaming(fileSet, all, again, o, w, r, handle.UnitFiles())
	}
}
