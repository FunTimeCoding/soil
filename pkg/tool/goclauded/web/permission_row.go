package web

import "github.com/funtimecoding/soil/pkg/tool/goclauded/permission"

func permissionRow(
	r *permission.Report,
	blocker string,
) *permission.Row {
	for _, w := range r.Rows {
		if w.Blocker == blocker {
			return w
		}
	}

	return nil
}
