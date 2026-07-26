package store

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"path/filepath"
	"strings"
)

func isHidden(path string) bool {
	for _, segment := range strings.Split(path, string(filepath.Separator)) {
		if strings.HasPrefix(segment, constant.Dot) {
			return true
		}
	}

	return false
}
