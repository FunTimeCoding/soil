package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunDefault(
	o *option.Fix,
	r *output.Results,
) {
	w := workspace.New(o.Diff, append([]string{o.Root}, o.Replacing...)...)
	w.BeforeCommit(o.BeforeCommit)
	runFix(o, w, r)
	parameterThrough(o.Patterns, o.Root, w, r)
	aliasThrough(o.Patterns, o.Root, w, r)
	formatThrough(o.Patterns, o.Root, w, r)
	finish(w, r)
}
