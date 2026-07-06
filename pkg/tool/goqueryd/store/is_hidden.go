package store

import (
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"path/filepath"
	"strings"
)

func isHidden(path string) bool {
	for _, segment := range strings.Split(path, string(filepath.Separator)) {
		if strings.HasPrefix(segment, separator.Dot) {
			return true
		}
	}

	return false
}
