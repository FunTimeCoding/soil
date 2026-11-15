package python_replacement

import "github.com/funtimecoding/soil/pkg/tool/goclaude/guard/shell_script"

func restored(
	r *Replacement,
	calls []*shell_script.Call,
) bool {
	if r.Target == "" {
		return false
	}

	for _, c := range calls {
		if c.Expr.Pos().Offset() > r.End && restores(c, r.Target) {
			return true
		}
	}

	return false
}
