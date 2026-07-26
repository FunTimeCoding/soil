package service

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func sessionSlug(path string) string {
	if i := strings.LastIndex(path, constant.Slash); i >= 0 {
		return path[:i]
	}

	return path
}
