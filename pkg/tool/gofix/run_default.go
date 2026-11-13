package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix/option"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
)

func RunDefault(
	o *option.Fix,
	r *output.Results,
) (w *workspace.Workspace, e error) {
	w = workspaceFor(o)
	defer func() { e = recoveredLoadFailure(recover()) }()
	w.BeforeCommit(o.BeforeCommit)

	if o.Loaded != nil && !o.Full {
		w.Remember(loadKey(o.Root, o.Patterns), o.Loaded, o.FileSet)
	}

	runFix(o, w, r)
	parameterThrough(o.Patterns, o.Root, w, r)
	aliasThrough(
		o.Patterns,
		o.Root,
		newAliasMode(false, defaultApplies, o.Ours),
		w,
		r,
	)
	formatThrough(o.Patterns, o.Root, w, r)
	finish(w, r)

	return w, nil
}
