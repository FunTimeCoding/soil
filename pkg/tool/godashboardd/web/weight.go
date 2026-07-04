package web

import "github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"

func weight(v *board.Section) int {
	result := len(v.Entries)

	if v.Name != "" {
		result++
	}

	return result
}
