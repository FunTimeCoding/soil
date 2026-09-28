package web

import "github.com/funtimecoding/soil/pkg/strings/join"

func truncate(
	v string,
	length int,
) string {
	if len(v) <= length {
		return v
	}

	return join.Empty(v[:length], "...")
}
