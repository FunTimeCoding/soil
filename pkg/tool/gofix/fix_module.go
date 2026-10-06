package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func fixModule(
	o *option.Fix,
	patterns []string,
	reported map[string]bool,
	faces *face.Set,
	w *workspace.Workspace,
	r *output.Results,
) {
	all, fileSet := loadThrough(o.Root, resolve.WithMainModule(patterns), w)

	if faces == nil {
		faces = face.New(all)
	}

	violations := FindViolations(all, reported, faces)

	if len(violations) > 0 {
		applyNaming(fileSet, all, violations, o, w, r, nil)
	}
}
