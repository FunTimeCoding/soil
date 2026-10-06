package web

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"

func weight(v *layout.Section) int {
	result := len(v.Entries)

	if v.Name != "" {
		result++
	}

	return result
}
