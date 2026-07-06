package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func shortenPath(path string) string {
	parts := split.Slash(path)

	if len(parts) <= 2 {
		return path
	}

	return join.Slash(parts[len(parts)-2:])
}
