package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard/edit"
	"mvdan.cc/sh/v3/syntax"
)

func restored(
	d *edit.Edit,
	calls []*syntax.CallExpr,
) bool {
	if d.Target == "" {
		return false
	}

	for _, c := range calls {
		if c.Pos().Offset() > d.End && restores(c, d.Target) {
			return true
		}
	}

	return false
}
