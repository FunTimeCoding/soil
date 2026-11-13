package gofix

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func runFix(
	o *option.Fix,
	w *workspace.Workspace,
	r *output.Results,
) {
	patterns := o.Patterns

	if len(patterns) == 0 {
		patterns = []string{goModConstant.AllPackages}
	}

	reported := o.Reported

	if reported == nil {
		reported = listReported(o.Root, patterns)
	}

	if o.Full || resolve.CoversMainModule(patterns) {
		fixModule(o, patterns, reported, nil, w, r)

		return
	}

	loaded, fileSet := loadThrough(o.Root, patterns, w)
	handle := indexHandle(o)
	faces := face.FromWorkspace(handle, loaded)
	violations := FindViolations(loaded, reported, faces)

	if len(violations) == 0 {
		return
	}

	if anyExported(violations) {
		targetedFix(o, patterns, reported, faces, violations, handle, w, r)

		return
	}

	applyNaming(fileSet, loaded, violations, o, w, r, nil)
}
