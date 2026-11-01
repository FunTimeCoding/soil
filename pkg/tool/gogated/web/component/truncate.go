package component

import "github.com/funtimecoding/soil/pkg/strings/join"

func Truncate(
	v string,
	length int,
) string {
	if len(v) <= length {
		return v
	}

	return join.Empty(v[:length], "...")
}
