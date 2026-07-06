package service

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func sessionSlug(path string) string {
	if i := strings.LastIndex(path, separator.Slash); i >= 0 {
		return path[:i]
	}

	return path
}
