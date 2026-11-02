package service

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func sessionSlug(path string) string {
	if i := strings.LastIndex(path, stringsConstant.Slash); i >= 0 {
		return path[:i]
	}

	return path
}
